# OpenAI Workshop

## The Goal of this Workshop

My goal is to give participants some ideas on what is possible to do using the
openAI API. Hopefully inspire new uses.

The driver for the learning will be to create a git commit suggestion tool.

## What do You Need

* Laptop
* Git repository to test with

**ADD SUGGESTION**

* OpenAI API key

## Workshop Structure

- Introduction 5 minutes
- Coding session 15 minutes
- OpenAI Talk 10 minutes
- Coding session 10 minutes
- Suggested improvements 5 minutes
- Coding session 10 minutes
- Wrap up & thoughts, 5 minutes


## Resources

- OpenAI billing limits: https://platform.openai.com/account/billing/limits
    It is recommended to use this to set a hard limit on your API usage.

- OpenAI playground https://platform.openai.com/playground/p/?mode=chat

- OpenAI Cookbook: https://github.com/openai/openai-cookbook/blob/main/techniques_to_improve_reliability.md
-
## What Is OpenAI?

OpenAI is a artificial intelligence research lab. 

The mission is to ensure that artificial general intelligence (AGI) benefits
all of humanity.

The organization started and continues to be a non-profit, but has a subsidiary
that is a capped for-profit company [1]. The goal with the structure is to try
to strike a balance between the slow pace of non-profit and the speed of a
for-profit.

## Models

The different models has different purposes, such as generating code, images,
natural language, etc. [3].

Example:

* DALL E: Image generation from natural language
* Whisper: Speech recognition model
* Embeddings: Measures relatedness of text strings. Use cases are search,
    recommendations, anomaly detection [4] etc.

## Note on Proprietary Data

Be mindful of the data you send into the API. Even though the data sent via the
API is not used to improve the models [5], you should have permission from your
client if you use any tools (such as the one we are building) professionally.

**FIXME**: Perhaps add some examples (amazon & samsung) where this has gone
wrong, though I think those were made before openAI changed the rule where the
API isn't used to improve the model.

## Cost Considerations

* Different models have different price points
* The pricing is pay as you go, and you pay per 1 000 /tokens/. 
* New accounts gets 5 dollar in free credit that can be used during your first 3
    months.

* Setup billing limits: https://platform.openai.com/account/billing/limits

See more [here](https://openai.com/pricing).

## Libraries

Libraries exists for many programming languages. See list
[here](https://platform.openai.com/docs/libraries).

It is also possible to write your own library if you wish.

## The tool

**SHOW DEMO**

* The file name is passed as a parameter.
* The git diff (file-content) is then parsed
* The git diff is then passed into a prompt sent to the OpenAI API.
* The response is returned from the CLI tool

## ⭐ Get Coding ⭐

15 minutes.

1. Get an OPENAI KEY. If you do now wish to create an account, send me a slack
   message and I'll send you one.
2. Decide on a strategy for how to interact with the tool you are creating
    2.1 Read filename
    2.2 Read entire diff
    2.3 Get diff using a git library
3. Get the diff into your tool
4. Send a request from your tool to the OpenAI API.
    Use a [library](https://platform.openai.com/docs/libraries) or build your
    own.

## ⏸️⏸️ Presentation

### OpenAI API Overview

#### Authentication

Authentication is done through bearer authentication (also called token
authentication). If you use a library the header might be set for you, but
you still need to provide it.

```sh
Authorization: Bearer OPENAI_API_KEY
```

#### Example Request

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

#### Data In Detail

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

- `models`: (required) specifies the models used. We will be using 3.5-turbo as it is the
    most capable.
    You can programatically get a list of models by doing a `GET` request to
    `https://api.openai.com/v1/models`.

- `messages`: (required) An array of messages that describes the conversation. 
      The role can be either `system`, `user` or `assistant`.

      As an example, this is one of the massages used for ChatGPT:

```sh
You are ChatGPT, a large language model trained by OpenAI. Answer as concisely
as possible. Knowledge cutoff: {knowledge_cutoff} Current date: {current_date}
```

      The system role is used to set the behaviour of the assistant.
      gpt3.5-turbo has a limitation on where it does not always pay strong
      attention to the system messages. In my experience a good system message
      together with an example user and assistant provides a good result.

      Another strategy is to only use a `user` message.

**ADD EXAMPLE**

- `temperature`: (optional) A value between 0 and 2 that decides how deterministic the
    response should be. 0 will be very deterministic (although not 100%
    deterministic) and a value with 2 will return more diverse completions.
    Default is 1.


#### Tokens

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

#### Limitations

- Knowledge cutoff

- Hallucinations

    Unfortunately the GPT model do not know the boundary of its knowledge very
    well. For the tool we are building it can have an effect where it will
    describe changes that are not there.

### Tips on git commands

Provide examples of git diffs with already made git commit messages to verify
against.

**FIXME**

#### Prompt Engineering

**FIXME**


## References

* [1]: https://youtu.be/L_Guz73e6fw?t=4434
     Sam Altman: OpenAI CEO on GPT-4, ChatGPT, and the Future of AI | Lex Fridman Podcast #367 @ youtube
* [2]: https://openai.com/blog/openai-lp
    OpenAI LP blog announcement @ openai.com
* [3]: https://platform.openai.com/docs/models/overview
    List of OpenAI API models
* [4]: https://platform.openai.com/docs/guides/embeddings/what-are-embeddings
    What are embeddings
* [5]: https://platform.openai.com/docs/guides/chat/do-you-store-the-data-that-is-passed-into-the-api
    FAQ: Do you store data that is passed into the API?
