# OpenAI Workshop

This is accompanying information for a mini workshop at tretton37.

## The Goal of this Workshop

The primary aim of this workshop is to spark creativity by demonstrating
the potential applications of the OpenAI API. 

We will drive this learning journey by focusing on the development of a Git
commit suggestion tool, showcasing practical implementation of the OpenAI API.

## Prerequisites

- Git
- OpenAI account

## Workshop Structure

- Introduction
- Coding Session 1
- OpenAI Details
- Coding Session 2
- Wrap Up 

## Introduction

An overview of OpenAI and its various AI models, emphasizes caution with
proprietary data, explains the pay-as-you-go pricing model, mentions available
programming libraries, and plans to detail strategies for interacting with the
tool. 

👉 [Introduction](./introduction.md) 👈

## Coding Session 1 

Let's get the ball rolling by getting an OpenAI key, devise a strategy for
integrating git diffs into their tool, and enable the tool to send requests to
the OpenAI API

👉 [Coding Session 1](./first-coding-session.md) 👈

## OpenAI API Overview

A quick insight on how to interact with the OpenAI API. Covers 'Authorization',
the chosen model, and a series of messages in the JSON payload. We explore
important parameters like 'temperature' that can influence the output's
randomness, and touch on the concept of tokens - central to usage and cost
implications.

Lastly, we delve into unique characteristics of the models including their
knowledge cutoff, potential for fabricating information, and their imperfect
word counting abilities. Strategies for effective git diff and commit as well as
tips for formulating compelling prompts also form part of the tutorial,
enhancing your overall proficiency with OpenAI's GPT-4 API.

👉 [OpenAI Overview](./openai-overview.md) 👈

## Coding Session 2

Lets get going with creating prompts and using this in our tool! After this you
should have a finished CLI tool to create commit message suggestions!

👉 [Coding Session 2](./second-coding-session.md) 👈

## Wrap Up

Proposed improvements, questions and further resources.

👉 [Wrap Up](./wrap-up.md) 👈

## Further Resources

- [ChatGPT Prompt Engineering for Developers](https://www.deeplearning.ai/short-courses/chatgpt-prompt-engineering-for-developers/)
- OpenAI playground https://platform.openai.com/playground/p/?mode=chat

- OpenAI Cookbook: https://github.com/openai/openai-cookbook/blob/main/techniques_to_improve_reliability.md

Examples are available in the [examples](./examples/) folder.
