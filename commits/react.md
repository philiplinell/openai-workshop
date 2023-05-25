# github.com/facebook/react

## commit message

```sh
commit f8de255e94540f9018d8196b3a34da500707c39b (HEAD)
Author: Andrew Clark <git@andrewclark.io>
Date:   Tue May 16 15:59:15 2023 -0400

    Lower Suspense throttling heuristic to 300ms (#26803)

    Now that the throttling mechanism applies more often, we've decided to
    lower this a tad to ensure it's not noticeable. The idea is it should be
    just large enough to prevent jank when lots of different parts of the UI
    load in rapid succession, but not large enough to make the UI feel
    sluggish. There's no perfect number, it's just a heuristic.
```

## diff

```sh
diff --git a/packages/react-reconciler/src/ReactFiberWorkLoop.js b/packages/react-reconciler/src/ReactFiberWorkLoop.js
index f6d1d7f7a..c558fbd21 100644
--- a/packages/react-reconciler/src/ReactFiberWorkLoop.js
+++ b/packages/react-reconciler/src/ReactFiberWorkLoop.js
@@ -375,7 +375,7 @@ let workInProgressRootRecoverableErrors: Array<CapturedValue<mixed>> | null =
 // content as it streams in, to minimize jank.
 // TODO: Think of a better name for this variable?
 let globalMostRecentFallbackTime: number = 0;
-const FALLBACK_THROTTLE_MS: number = 500;
+const FALLBACK_THROTTLE_MS: number = 300;
 
 // The absolute time for when we should start giving up on rendering
 // more and prefer CPU suspense heuristics instead.
diff --git a/packages/react-reconciler/src/__tests__/ReactSuspenseWithNoopRenderer-test.js b/packages/react-reconciler/src/__tests__/ReactSuspenseWithNoopRenderer-test.js
index fc1aa3870..1b05f8a2f 100644
--- a/packages/react-reconciler/src/__tests__/ReactSuspenseWithNoopRenderer-test.js
+++ b/packages/react-reconciler/src/__tests__/ReactSuspenseWithNoopRenderer-test.js
@@ -1863,8 +1863,8 @@ describe('ReactSuspenseWithNoopRenderer', () => {
     // Advance by a small amount of time. For testing purposes, this is meant
     // to be just under the throttling interval. It's a heurstic, though, so
     // if we adjust the heuristic we might have to update this test, too.
-    Scheduler.unstable_advanceTime(400);
-    jest.advanceTimersByTime(400);
+    Scheduler.unstable_advanceTime(200);
+    jest.advanceTimersByTime(200);
 
     // Now resolve B.
     await act(async () => {
```
