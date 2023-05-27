# Second Coding Session

~ 20 minutes.

Goals to achieve:

1. Create a prompt.

I recommend using [ChatGPT](https://chat.openai.com/) with the prompt described
[here](https://github.com/philiplinell/openai-workshop/blob/main/openai-overview.md#prompt-engineering)
to generate the prompt.

For testing I recommend you to use the
[playground](https://platform.openai.com/playground/p/?mode=chat).

For example git diffs see 👉 [here](./commits/README.md) 👈
Useful when testing your prompt!

2. Use the prompt in your CLI tool

3. (Optional) Set it up as a git hook so that it will run on each git commit!

4. Improvements?!

## 1. Create a Prompt

You can iterate yourself on [playground](https://platform.openai.com/playground/p/?mode=chat).

Or you can use [chat.openai.com](https://chat.openai.com/) and use the following
prompt:

```markdown
As my Prompt Architect, I'm seeking your assistance in developing an optimal prompt for my needs to be used with the OpenAI API. Our collaboration will follow an iterative process, as detailed below:

1. **Topic Identification:** Start by inquiring about the intended topic of the prompt. I will provide an initial idea, which will set the stage for our iterative refinement process. 

2. Based on my input, you will then elaborate on three areas:
    a) **Prompt Refinement:** In this section, you will present a revised version of my initial prompt, aiming for precision, conciseness, and easy understanding. 
    b) **Detail Extraction:** Here, you should propose questions that could extract additional information or specific details from me, aiding the prompt's further refinement. 
    c) **Feedback Solicitation:** Request my opinion on the refined prompt. Inquire if there are any aspects of the prompt that require adjustment or if I have feedback that could further improve it.

3. Our collaboration will advance in this iterative manner - I will offer more information, share my feedback, and you will continuously refine the prompt under the 'Prompt Refinement' section. This cycle will continue until I confirm that the prompt meets my expectations.

Please remember, your suggestions and questions should consistently aim to improve the prompt's focus, minimize ambiguity, and enhance its effectiveness. Let's embark on this process!
```

## 2. Use the prompt in your CLI tool


Example snippet (node)

```javascript

// ...

// Function to generate the prompt for the AI model
function createPrompt() {
    return "Given the following git diff, which contains the lines changed and filenames, please provide an appropriate commit message suggestion. Make sure ..."; // intentionally cut of. Use your own crafted prompt!
}

// Main function
async function main() {
    try {
        // Reading the file's content
        const gitDiff = fs.readFileSync(filename, 'utf8');
        // Creating the AI prompt
        const prompt = createPrompt();
        // Combining the prompt and the file's content
        const content = `${prompt}\n\n${gitDiff}`;

        // Sending a request to the OpenAI API and waiting for the result
        const completion = await openai.createChatCompletion({
            model: "gpt-3.5-turbo",
            messages: [{role: "user", content: content}],
        });

        // Logging the AI's response
        console.log(completion.data.choices[0].message.content);

        // ...

```

## 3. Use it together with a git hook!


This is an example on how to do just that (where the CLI tool is called
`commit-msg` and is in `PATH`):

```sh
───────┬────────────────────────────────────────────────────────────────────────────────────
       │ File: .git/hooks/prepare-commit-msg
───────┼────────────────────────────────────────────────────────────────────────────────────
   1   │ #!/bin/sh
   2   │
   3   │ # Use CLI tool commit-msg to fetch a suggested commit message. Prepend the
   4   │ # suggested commit message to the commit message file.
   5   │
   6   │ COMMIT_MSG_FILE=$1
   7   │
   8   │ echo "Fetching suggested commit message..."
   9   │
  10   │ COMMIT_MSG=$(commit-msg --timeout=15s --file=$COMMIT_MSG_FILE)
  11   │
  12   │ if [ $? -ne 0 ]; then
  13   │     echo "❌ prepare-commit-msg: commit-msg failed. Doing nothing..."
  14   │     exit 0
  15   │ fi
  16   │
  17   │ printf '%s\n%s\n' "${COMMIT_MSG}" "$(cat $COMMIT_MSG_FILE)" >$COMMIT_MSG_FILE
───────┴────────────────────────────────────────────────────────────────────────────────────
```

## 4. Improvements

Some suggested improvements are [here](./wrap-up.md)
