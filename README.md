[![Build Status](https://drone.io/github.com/fblecha/fuel/status.png)](https://drone.io/github.com/fblecha/fuel/latest)
[![GoDoc](https://godoc.org/github.com/fblecha/fuel?status.svg)](https://godoc.org/github.com/fblecha/fuel)


# Purpose

Fuel is a static website creator -- it takes in markdown files and produces a static html site you can host from [github pages](https://pages.github.com/), [Amazon S3](https://aws.amazon.com/s3/), or any web host.  Several tools already support this, notably [Jekyl](https://jekyllrb.com/) and [Hugo](https://gohugo.io/).  

## So what makes Fuel different?  
Fuel also has the ability to take in JSON-enhanced markdown (I should totally trademark that) and:
1. use the JSON in your website **as variables**
2. store the JSON **and** content into a database   

# How does it work?

Given a normal markdown file or (as in the example below) a JSON-enahanced markdown file:

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

3. Complete the [go installation instructions](https://golang.org/doc/install).
4. Get fuel

```shell
$ go get github.com/fblecha/fuel
$ cd $GOPATH/src/github.com/fblecha/fuel
```
5. Install fuel

This will install fuel into your $GOPATH/bin directory.
```shell
$ go install
```


# How do you run it?
Fuel is a command-line application, so you'll run it via a terminal window.

If you want to create a blog devoted to dogs, then you would do the following steps:

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

## What happens when fuel is run?
Fuel does the following:

1. It creates a public directory.  This directory will be the home of your website.
2. it copies everything from ./style over to ./public/style.  If you want to copy over CSS, JavaScript, or anything else, put it in your style directory and refer to it as "/style/someting" in your html files.
3. It looks for each markdown file under ./content and applies a layout (a HTML file) to it so that you end up with ./public/something.html

## How do layouts work?

Insert


# What databases are supported?
* (Not started) postgresql
* (Not started) DynamoDB


# (Developers) What else needs to be done?

see the [Todo file](./Todo.md)
