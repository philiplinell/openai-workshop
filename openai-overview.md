# OpenAI API Overview

We will be using the [Chat
completion](https://platform.openai.com/docs/guides/chat) API.

## Authentication

Authentication is done through bearer authentication (also called token
authentication). If you use a library the header might be set for you, but
you still need to provide it.

```sh
Authorization: Bearer OPENAI_API_KEY
```

### Example Request

```sh
curl https://api.openai.com/v1/chat/completions \
  --header "Content-Type: application/json" \
  --header "Authorization: Bearer $OPENAI_API_KEY" \
  --data '{
     "model": "gpt-3.5-turbo",
     "messages": [{"role": "user", "content": "Say this is a test!"}],
     "temperature": 0.7
   }'
```

### Data In Detail

```json
{
 "model": "gpt-3.5-turbo",
 "messages": [
    {"role": "system", "content": "You are a helpful assistant."},
    {"role": "user", "content": "Who won the world series in 2020?"},
    {"role": "assistant", "content": "The Los Angeles Dodgers won the World Series in 2020."},
    {"role": "user", "content": "Where was it played?"}
 ],
 "temperature": 0.7
}
```

- `model`: (required) specifies the model used. We will be using 3.5-turbo as it is the
    most capable.
    You can programmatically get a list of models by doing a `GET` request to
    `https://api.openai.com/v1/models`.

- `temperature`: (optional) A value between 0 and 2 that decides how deterministic the
    response should be. 0 will be very deterministic (although not 100%
    deterministic) and a value with 2 will return more diverse completions.
    Default is 1.

- `messages`: (required) An array of messages that describes the conversation. 
      The role can be either `system`, `user` or `assistant`.

The `user` and `assistant` roles can be used both when creating a chat bot where
you want to include the previous messages (as a history) or for initiating new
conversations.

If you are using it for historical purposes you include past `user` and
`assistant` messages. This will enable the model to consider the entire
conversation history when generating a response. Currently there is no way to
keep state across API requests as each API request has to include previous
conversations if you want to continue a conversation.
Note though that the entire message can be too long for the API call. In that
case you need to shorten it somehow.

Initiating new conversations: When starting a new conversation, you can use the
"user" role to pose the initial question or statement, and the "assistant" role
to provide the initial response.


As an example, this is one of the messages used for ChatGPT:

```sh
You are ChatGPT, a large language model trained by OpenAI. Answer as concisely
as possible. Knowledge cutoff: {knowledge_cutoff} Current date: {current_date}
```

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

## Tokens

OpenAIs models processes text by breaking them down into units called tokens.

A token is, roughly, 0.75 word but longer words will be more than one token.

"Common words like “cat” are a single token, while less common words are often
broken down into multiple tokens. For example, “Butterscotch” translates to four
tokens: “But”, “ters”, “cot”, and “ch”. "

Each model has a maximum token count. This is counting both the request and
response. GPT 3.5 turbo has a maximum limit of 4096 token (roughly 3000 words).

E.g. if your API call is 10 tokens in the message and you recieve 20 tokens in
the output, you will be billed for 30 tokens.

The response will contain the `total_tokens` used.

## Limitations

- Knowledge cutoff

    The knowledge cut off date is September 2021. This means that the model can
    generate outdated information and code. 

- Hallucinations

    Unfortunately the GPT model do not know the boundary of its knowledge very
    well. It can give incomplete or wrong answers, and will do so with
    confidence.

    The fabricated ideas are called *hallucinations*.

    For the tool we are building it can have an effect where it will
    describe changes that are not there.

- Counting words

    Large language models are not good at returning a specific word count.
    Instead, if you wish to limit the response use prompts such as "use at most
    3 sentences" or "use at most 320 characters".

## Prompt Engineering

The main principle is to use clear & specific instructions, but make sure to
distinguish writing a clear prompt from writing a short one. A longer prompt
provides more clarity and context for the model, leading to more detailed and
relevant outputs. With that said, the prompt will count toward the tokens used,
so a trade-off has to be made where the prompt is clear-and-specific enough.

The perfect prompt is **rarely created on the first try**. Instead try an
iterative process where the prompt is refined.

My favorite strategy to create a good prompt is to utilize ChatGPT with a prompt
that develops another prompt <insert Xzibit Yo Dawg Meme>:

```
As my Prompt Architect, I would like your assistance in developing the most effective prompt for my needs, which will be utilized by you, ChatGPT. Follow the steps below to ensure a collaborative and iterative process:

1. Initiate by asking about the desired topic of the prompt. I will provide an initial response which will serve as the foundation for our iterative refinement process. 

2. Based on my input, you will generate three sections:
    a) **Prompt Enhancement:** Here, you should present your revised version of my original prompt, aiming for clarity, brevity, and comprehensibility. 
    b) **Inquiry for Details:** Pose any relevant questions here that could help gather more specific information or details from me to further refine the prompt. 
    c) **Feedback and Adjustments:** Ask me if there are any areas in the revised prompt that need adjustment or if I have any feedback to provide.

3. We will progress in this iterative manner with me supplying more details, providing feedback, and you continually fine-tuning the prompt under the 'Prompt Enhancement' section. This process will repeat until I confirm that we have achieved the desired prompt.

Please note, your inputs and questions should always be designed to help sharpen the focus, reduce ambiguity, and increase the effectiveness of the prompt. Let's begin!
```

And for the API (which is very similar to above but the usage is changed):

```
As my Prompt Architect, I'm seeking your assistance in developing an optimal prompt for my needs to be used with the OpenAI API. Our collaboration will follow an iterative process, as detailed below:

1. **Topic Identification:** Start by inquiring about the intended topic of the prompt. I will provide an initial idea, which will set the stage for our iterative refinement process. 

2. Based on my input, you will then elaborate on three areas:
    a) **Prompt Refinement:** In this section, you will present a revised version of my initial prompt, aiming for precision, conciseness, and easy understanding. 
    b) **Detail Extraction:** Here, you should propose questions that could extract additional information or specific details from me, aiding the prompt's further refinement. 
    c) **Feedback Solicitation:** Request my opinion on the refined prompt. Inquire if there are any aspects of the prompt that require adjustment or if I have feedback that could further improve it.

3. Our collaboration will advance in this iterative manner - I will offer more information, share my feedback, and you will continuously refine the prompt under the 'Prompt Refinement' section. This cycle will continue until I confirm that the prompt meets my expectations.

Please remember, your suggestions and questions should consistently aim to improve the prompt's focus, minimize ambiguity, and enhance its effectiveness. Let's embark on this process!
```

### Techniques

Here are some techniques you can use while refining your prompt.

#### Use delimiters to clearly indicate distinct parts.

Delimiters can be triple backticks, triple quotes """, triple dashes, angled
brackets, `xml` tags, etc.

It is important to note that GPT-3.5 doesn't treat backticks as special or
distinguishing in any way. They are treated as ordinary characters and don't
serve the function they do in Markdown or other such languages.

Choose a delimiter that is unlikely to be part of the user-generated input.

Delimiters are also helpful in avoiding prompt injection. Prompt injection is when a
user is allowed to add some input to your prompt and could potentially give
conflicting instructions to the model.

#### Structured output

Ask for a response in a specific format which can make the model response easier
to response.

E.g.

message (role `user`)
```sh
Generate 3 made-up planets for a sci-fi book along with planet characteristics. 
Provide them in JSON format with the following keys: in_habitable_zone (bool),
atmospheric_composition (string), average_temperature (float).
Only respond with the JSON.
```

Response:

```json
{
  "planet_1": {
    "in_habitable_zone": true,
    "atmospheric_composition": "nitrogen, oxygen, carbon dioxide",
    "average_temperature": 25.5
  },
  "planet_2": {
    "in_habitable_zone": false,
    "atmospheric_composition": "methane, ammonia, hydrogen",
    "average_temperature": -150.2
  },
  "planet_3": {
    "in_habitable_zone": true,
    "atmospheric_composition": "helium, neon, argon",
    "average_temperature": -80.9
  }
}
```


#### Give the model an out

Specify what the model should do in case that any preconditions are not met.

```
Summarize the provided receipe and re-write it as clear instructions in the
following format:

Step 1: ...
Step 2: ...
...
Step N: ...

If the text do not contain a recepie, then simple write "No recepie provided".

${recepie}
```


## OpenAI Playground

OpenAI Playground https://platform.openai.com/playground/p/?mode=chat
