# github.com/angular/angular 

## commit message

```sh
commit 2cdb4c5911965aa273f11432e04502e52b5e1b9b (HEAD)
Author: Alan Agius <alanagius@google.com>
Date:   Tue May 23 14:28:46 2023 +0000

    fix(http): create macrotask during request handling instead of load start (#50406)

    This commit schedules the macrotask creation to happen before the XHR `loadStart` event. This is needed as in some cases, Zone.js becomes stable too early.

    With this commit, we also update the internal `createBackgroundMacroTask` method to use Zone.js `scheduleMacroTask` as otherwise the `setTimeout` would cause `fakeAsync` tests to fail due to pending timers.

    Closes #50405

    PR Close #50406
```

## diff

```sh
diff --git a/packages/common/http/src/xhr.ts b/packages/common/http/src/xhr.ts
index 8a6e2300f8..ce24a22013 100644
--- a/packages/common/http/src/xhr.ts
+++ b/packages/common/http/src/xhr.ts
@@ -60,6 +60,10 @@ export class HttpXhrBackend implements HttpBackend {
               `Cannot make a JSONP request without JSONP support. To fix the problem, either add the \`withJsonpSupport()\` call (if \`provideHttpClient()\` is used) or import the \`HttpClientJsonpModule\` in the root NgModule.`);
     }
 
+    // Schedule a macrotask. This will cause NgZone.isStable to be set to false,
+    // Which delays server rendering until the request is completed.
+    const macroTaskCanceller = createBackgroundMacroTask();
+
     // Check whether this factory has a special function to load an XHR implementation
     // for various non-browser environments. We currently limit it to only `ServerXhr`
     // class, which needs to load an XHR implementation.
@@ -231,7 +235,11 @@ export class HttpXhrBackend implements HttpBackend {
                 statusText: xhr.statusText || 'Unknown Error',
                 url: url || undefined,
               });
+
               observer.error(res);
+
+              // Cancel the background macrotask.
+              macroTaskCanceller();
             };
 
             // The sentHeaders flag tracks whether the HttpResponseHeaders event
@@ -309,35 +317,31 @@ export class HttpXhrBackend implements HttpBackend {
               }
             }
 
-            let macroTaskCanceller: VoidFunction|undefined;
-
             /** Tear down logic to cancel the backround macrotask. */
-            const onLoadStart = () => {
-              macroTaskCanceller ??= createBackgroundMacroTask();
-            };
-            const onLoadEnd = () => {
-              macroTaskCanceller?.();
-            };
+            const onLoadEnd = () => macroTaskCanceller();
 
-            xhr.addEventListener('loadstart', onLoadStart);
             xhr.addEventListener('loadend', onLoadEnd);
 
             // Fire the request, and notify the event stream that it was fired.
-            xhr.send(reqBody!);
+            try {
+              xhr.send(reqBody!);
+            } catch (e: any) {
+              onError(e);
+            }
+
             observer.next({type: HttpEventType.Sent});
             // This is the return from the Observable function, which is the
             // request cancellation handler.
             return () => {
               // On a cancellation, remove all registered event listeners.
-              xhr.removeEventListener('loadstart', onLoadStart);
               xhr.removeEventListener('loadend', onLoadEnd);
               xhr.removeEventListener('error', onError);
               xhr.removeEventListener('abort', onError);
               xhr.removeEventListener('load', onLoad);
               xhr.removeEventListener('timeout', onError);
 
-              //  Cancel the background macrotask.
-              macroTaskCanceller?.();
+              // Cancel the background macrotask.
+              macroTaskCanceller();
 
               if (req.reportProgress) {
                 xhr.removeEventListener('progress', onDownProgress);
@@ -357,11 +361,8 @@ export class HttpXhrBackend implements HttpBackend {
   }
 }
 
-// Cannot use `Number.MAX_VALUE` as it does not fit into a 32-bit signed integer.
-const MAX_INT = 2147483647;
-
 /**
- * A method that creates a background macrotask of up to Number.MAX_VALUE.
+ * A method that creates a background macrotask using Zone.js.
  *
  * This is so that Zone.js can intercept HTTP calls, this is important for server rendering,
  * as the application is only rendered once the application is stabilized, meaning there are pending
@@ -370,7 +371,15 @@ const MAX_INT = 2147483647;
  * @returns a callback method to cancel the macrotask.
  */
 function createBackgroundMacroTask(): VoidFunction {
-  const timeout = setTimeout(() => void 0, MAX_INT);
+  // We use Zone.js when it's defined as otherwise a `setTimeout` will leave open timers which
+  // causes `fakeAsync` tests to fail.
+  const noop = () => {};
+  if (typeof Zone !== 'undefined') {
+    const zoneCurrent = Zone.current;
+    const task = zoneCurrent.scheduleMacroTask('httpMacroTask', noop, undefined, noop, noop);
+
+    return () => zoneCurrent.cancelTask(task);
+  }
 
-  return () => clearTimeout(timeout);
+  return noop;
 }

```
