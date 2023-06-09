# Introduction

## Disclaimer

I'm no expert. I'm basing the information here from personal experience and from
what I've learnt online.

## What Will We Build?

**SHOW DEMO**

* The file name is passed as a parameter.
* The git diff (file-content) is then parsed
* The git diff is then passed into a prompt sent to the OpenAI API.
* The response is returned from the CLI tool

Commands used in demo:
```sh
# Show command

# Angular demo
cd angular
git reset HEAD~
git add .
git commit
```

## What Is OpenAI?

OpenAI is a artificial intelligence research lab. 

The mission is to ensure that artificial general intelligence (AGI) benefits
all of humanity.

The organization started and continues to be a non-profit, but has a subsidiary
that is a capped for-profit company [^1] [^2]. The goal with this structure is to try
to strike a balance between the slow pace of non-profit and the speed of a
for-profit.

## Models

The different models has different purposes, such as generating code, images,
natural language, etc. [^3].

| Model | Purpose |
| -- | -- |
| DALL E | Image generation from natural language |
| Whisper | Speech recognition model |
| Embeddings |  Measures relatedness of text strings. Use cases are search, recommendations, anomaly detection [^4] etc. |

For our project, we are utilizing **GPT-3**, a language model, due to its ability
to understand and generate meaningful responses from text inputs.

## Note on Proprietary Data

Be mindful of the data you send into the API. Even though the data sent via the
API is not used to improve the models [^5], you should have permission from your
client if you use any tools (such as the one we are building) professionally.

For example, in the past, companies like Amazon and Samsung faced issues related
to data privacy and misuse. These examples underscore the importance of
carefully managing the data you share with AI APIs.

## Cost Considerations

* Different models have different price points
* The pricing is pay as you go, and you pay per 1 000 *tokens*. 
* New accounts gets 5 dollar in free credit that can be used during your first 3
    months.

* Setup billing limits: https://platform.openai.com/account/billing/limits

See more on [openai.com/pricing](https://openai.com/pricing).

## Libraries

Libraries exist for many programming languages. See the list [openai.com/docs/libraries](https://platform.openai.com/docs/libraries).

It is also possible to write your own library if you wish.


## Strategy for interacting with our CLI tool

* Pass filename that contains git diff

Probably easiest ⭐

E.g.
```sh
$ commit-msg --filename=.git/COMMIT_EDITTMSG
```

* Read entire diff

E.g.
```sh
$ commit-msg --changes=$(git diff)
```

* Get diff using a git library

This would use a git library (or call git as a shell command) to get a git diff.

## References

[^1]: https://youtu.be/L_Guz73e6fw?t=4434
     Sam Altman: OpenAI CEO on GPT-4, ChatGPT, and the Future of AI | Lex Fridman Podcast #367 @ youtube

[^2]: https://openai.com/blog/openai-lp
    OpenAI LP blog announcement @ openai.com

[^3]: https://platform.openai.com/docs/models/overview
    List of OpenAI API models

[^4]: https://platform.openai.com/docs/guides/embeddings/what-are-embeddings
    What are embeddings

[^5]: https://platform.openai.com/docs/guides/chat/do-you-store-the-data-that-is-passed-into-the-api
    FAQ: Do you store data that is passed into the API?
