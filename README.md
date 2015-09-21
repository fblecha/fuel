[![Build Status](https://drone.io/github.com/fblecha/haiku/status.png)](https://drone.io/github.com/fblecha/haiku/latest)
[![GoDoc](https://godoc.org/github.com/fblecha/haiku?status.svg)](https://godoc.org/github.com/fblecha/haiku)
[![Coverage Status](https://coveralls.io/repos/fblecha/haiku/badge.svg?branch=master&service=github)](https://coveralls.io/github/fblecha/haiku?branch=master)


WIP - do not use until there's a release

# haiku
A static blog that can also persist data to a database.

# Purpose

I wanted to create a series of content that would be available by both an API and via a static website, and I didn't want to use wordpress.  Haiku will do two major things:
1. produce a static website
2. for each type of content, persist both the data and metadata into a database

## Databases supported
* (WIP) postgresql
* (Not started) DynamoDB
## Expected support
- Create new content
- Publish new content
- Save you content to a database

# How does it work?

Given a .haiku file, like this:

```markdown
{
  "breed": "Labrador Retriever"
  "colors": [
    "black",
    "yellow",
    "chocolate"
  "]
}
~~~
\# Breed: Labrador Retriever

The Labrador Retriever, also known as simply Labrador or Lab, is one of several kinds of retrievers, a type of [gun dog](https://en.wikipedia.org/wiki/Gun_dog). Labradors are athletic, playful, and the most popular breed of dog by registered ownership in Australia, Canada, New Zealand,[4] the United Kingdom,[5] and the United States (since 1991).[6]

[via Wikipedia](https://en.wikipedia.org/wiki/Labrador_Retriever)

```
