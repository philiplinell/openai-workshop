# Wrap Up

## Suggested Improvements

### Take your Commit messages into account

Instead of basing the commit message from only the git diff, why not create
prompt that accepts both an initial commit message from you and the git diff.

### Add Style and/or pizzazz

**Descriptive and Neutral Style:**

```
The style of the commit message should be descriptive and neutral. Use clear,
concise language to describe the changes. The message should be objective and
factual, focusing solely on what was done, without injecting personal opinions
or unnecessary context. It should provide enough information for a reader to
understand the changes without having to look at the code.
```

**List-Based Style:**

```
The style of the commit message should be list-based. Use bullet points or
numbered lists to itemize the changes made. Each point should be concise,
specific, and self-explanatory. This style is particularly suitable for commits
that involve multiple changes or updates. Despite the structured format, ensure
the message provides enough context to understand the changes without having to
look at the code.
```

**Problem-Solution Style:**

```
The style of the commit message should be problem-solution oriented. Begin by
clearly outlining the problem or issue that was addressed. Follow this with a
concise explanation of the solution implemented to fix the problem. This style
encourages a logical and methodical approach to describing changes, and is
particularly effective for commits aimed at fixing bugs or improving
functionality. Ensure the message provides enough context to understand the
changes without having to look at the code.
```

Less serious:

**Rap Song Style:**

```
The style of the commit message should be akin to a rap song. Craft a rhyming
couplet or a short verse that describes the changes. Use rhythm and rhyme to
make your message catchy and memorable. While having fun with this style, ensure
the message still provides enough context to understand the changes without
having to look at the code.
```

**Haiku Style:**

```
The style of the commit message should be similar to a haiku. Craft a three-line
poem with a 5-7-5 syllable count that encapsulates the essence of the changes.
This style encourages creativity and brevity. Despite the poetic nature, ensure
the message provides enough context to understand the changes without having to
look at the code.
```

**Famous Quote Style:**

```
The style of the commit message should mirror a famous quote, but adapted to
describe the changes made. Think of an inspiring or humorous quote, then modify
it to fit your commit. This style can add a layer of intrigue and wit to your
message. Despite the creative twist, ensure the message provides enough context
to understand the changes without having to look at the code.
```

### Integrate it into your favorite editor

Instead of triggering the tool by a Git hook (or whatever method you choose),
why not integrate the CLI tool into your favorite editor? This way it can be
called upon when you most need it.

### Handle too long git diffs

As both the prompt, git diff and response will count towards the max tokens you
might want to shorten the git diff if it is too long.

## Anyone wants to demo their tool?

## Questions

## Further Resources

- [ChatGPT Prompt Engineering for Developers](https://www.deeplearning.ai/short-courses/chatgpt-prompt-engineering-for-developers/)

- OpenAI playground https://platform.openai.com/playground/p/?mode=chat

- OpenAI Cookbook: https://github.com/openai/openai-cookbook/blob/main/techniques_to_improve_reliability.md

## Feedback link

Go to www.menti.com.

Enter code 41 18 90 9.

Or scan QR Code

![](./qr.png)
