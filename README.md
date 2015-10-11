[![Build Status](https://drone.io/github.com/fblecha/fuel/status.png)](https://drone.io/github.com/fblecha/fuel/latest)
[![GoDoc](https://godoc.org/github.com/fblecha/fuel?status.svg)](https://godoc.org/github.com/fblecha/fuel)


WIP - do not use until there's a release

# fuel
A app that will create a blog, and optionally store it in a database.

# Purpose

I wanted to create a series of content that would be available by both an API and as a static website.  fuel will do two major things:
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

Given a .fuel file, like this:

```markdown
{
  "breed": "Labrador Retriever"
  "colors": [
    "black",
    "yellow",
    "chocolate"
  "]
}
\~~~
# Breed: Labrador Retriever

The Labrador Retriever, also known as simply Labrador or Lab, is one of several
kinds of retrievers, a type of [gun dog](https://en.wikipedia.org/wiki/Gun_dog).
 Labradors are athletic, playful, and the most popular breed of dog by registered
 ownership in Australia, Canada, New Zealand,[4] the United Kingdom,[5] and
 the United States (since 1991).[6]

[via Wikipedia](https://en.wikipedia.org/wiki/Labrador_Retriever)

```

# How do you run it?
```shell
# to create a new blog
$ fuel my_new_dog_blog
$ cd my_new_dog_blog

# create some markdown files under my_new_dog_blog/content
# create some layout.html files under my_new_dog_blog/views
# optionally put any css/images/etc you need under my_new_dog_blog/style

$ fuel run

# load my_new_dog_blog/public in a web server
```
