# Second Coding Session

~ 20 minutes.

Goals to achieve:

1. Create a prompt.

I recommend using [ChatGPT](https://chat.openai.com/) with the prompt described
[here](https://github.com/philiplinell/openai-workshop/blob/main/README.md#prompt-engineering)
to generate the prompt.

For testing I recommend you to use the
[playground](https://platform.openai.com/playground/p/?mode=chat).

For example git diffs see 👉 [here](./commits/README.md) 👈
Useful when testing your prompt!

2. Use the prompt in your CLI tool

3. (Optional) Set it up as a git hook so that it will run on each git commit!

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

