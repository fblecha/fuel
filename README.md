[![Build Status](https://drone.io/github.com/fblecha/fuel/status.png)](https://drone.io/github.com/fblecha/fuel/latest)
[![GoDoc](https://godoc.org/github.com/fblecha/fuel?status.svg)](https://godoc.org/github.com/fblecha/fuel)


WIP - do not use until there's a release

# Purpose

I wanted to create a series of content that would be available by both an API and as a static website.  fuel will do two major things:
1. produce a static website
2. for each type of content, persist both the data and metadata into a database

## Databases supported
* (Not started) postgresql
* (Not started) DynamoDB

## Expected support
- Create new content
- Publish new content
- Save you content to a database

# How does it work?

Given a .md file, like this:

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
fuel will convert it into a HTML file.

# Installation

Type the following commands into your Terminal.

1. Install homebrew (see http://brew.sh/)
2. Install go
```shell
$ brew install go
```
3. Edit your $GOPATH

You can edit it manually view your terminal:
```shell
export GOPATH=$HOME/go
```
or preferably you can edit your .bashrc file (assuming you're using bash) to have it.

4. Get fuel

```shell
$ go get github.com/fblecha/fuel
$ cd ~/go/src/github.com/fblecha/fuel
```

5. Install fuel

This will install fuel into your $GOPATH/bin directory.
```shell
$ go install
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

## Todo

see the [Todo file](./Todo.md)
