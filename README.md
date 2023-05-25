- OpenAI Details 10 minutes
- Coding session 20 minutes
that is a capped for-profit company [^1]. The goal with this structure is to try
natural language, etc. [^2].
| Model | Purpose |
| -- | -- |
| DALL E | Image generation from natural language |
| Whisper | Speech recognition model |
| Embeddings |  Measures relatedness of text strings. Use cases are search, recommendations, anomaly detection [^3] etc. |
API is not used to improve the models [^4], you should have permission from your
* The pricing is pay as you go, and you pay per 1 000 *tokens*. 



As an example, this is one of the messages used for ChatGPT:
The system role is used to set the behaviour of the assistant.
gpt3.5-turbo has a limitation on where **it does not always pay strong
attention to the system messages** ⚠️.

There are a couple of strategies to work around this:

<details>

<summary>1. Use a system message with examples as user and assistant messages</summary>


The `user` and `assistant` roles will provide conversation examples to guide
the model's behaviour:

message (role `system`)

```sh
You are an insightful assistant that crafts
commit messages. The commit messages should accurately and succinctly explain
the changes made in the files, detailing the reason for changes and the effect
they will have on the project. Your responses should consist of the commit
subject and the commit body, separated by newlines.

The commit subject should:
- Be brief (50 characters or less)
- Use the imperative mood (e.g., "Add", "Fix", "Change")

The commit body should:
- Further explain the changes in detail if necessary
- Be wrapped at 72 characters
- Be separated from the commit subject by a blank line

Make sure they provide enough context to understand the changes without having to look at the code.

The style of the commit message should be conversational and casual using informal language or even a touch of humor to describe the changes. You should aim to make the commit messages engaging, yet still professional and informative.
```

message (role `user`)

This is used as an example for the model.

```git
diff --git a/README.md b/README.md
new file mode 100644
index 0000000..ca34b6a
--- /dev/null
+++ b/README.md
@@ -0,0 +1,21 @@
+# Commit Message
+
+Create a commit message suggestion from the git diff using the openAI API.
+
+Note that this means that filename and lines changed is sent to openAI. If that
+bothers you - don't use this tool.
```

message (role `assistant`)
This is used as an example response for the model.
```sh
Unleashing a brand new README.md to demystify our OpenAI-powered commit message wizardry!\n\nHey folks,\nWe just slapped a shiny new README.md into the mix! 🎉\nThis bad boy's job is to school you all about our super cool, freshly baked tool that spits out commit message suggestions - all powered by the magic of OpenAI (no wizards were harmed in the process, promise! 🧙.\nIt's got everything - the ins, the outs, the what-have-yous about our tool. Oh, and it's also gonna give you the lowdown on the stuff we're sending over to OpenAI (don't worry, it's just filenames and changed lines, not your secret cookie recipes! 🍪).\nSo strap in, take a gander at the README, and let's get those commit messages singing!
```

message (role `user`)

This is the final message that contains the git diff.

```git
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

</details>

<details>

<summary>2. Use only use a single `user` role message</summary>

This will use less tokens.

message (role `user`)
```git
Given the git diff below, which contains the lines changed and filenames, please
provide an appropriate commit message suggestion. Make sure to highlight any
breaking changes explicitly. The commit message should consist of a subject and
a body, separated by two newlines. The subject, written in the imperative mood
(e.g., "Add", "Fix", "Change"), should be brief, 50 characters or less. The body
of the message should be wrapped at 72 characters.

Git diff:

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
</details>
The perfect prompt is **rarely created on the first try**. Instead try an
message (role `user`)
```sh
Provide them in JSON format with the following keys: in_habitable_zone (bool),
atmospheric_composition (string), average_temperature (float).
```json
    I recommend using [ChatGPT](https://chat.openai.com/) with the prompt
    [above](https://github.com/philiplinell/openai-workshop/blob/main/README.md#prompt-engineering)
`commit-msg` and is in `PATH`):
[^1]: https://youtu.be/L_Guz73e6fw?t=4434

[^2]: https://openai.com/blog/openai-lp

[^3]: https://platform.openai.com/docs/models/overview

[^4]: https://platform.openai.com/docs/guides/embeddings/what-are-embeddings

[^5]: https://platform.openai.com/docs/guides/chat/do-you-store-the-data-that-is-passed-into-the-api