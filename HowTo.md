
For this example, let's assume you want to start a new blog on dog breeds and you want to put together a site that explains each breed.

Assuming you start in your home directory (for example ~ or /Users/me), then let's assume you want your new site in the ~/dogblog directory.

## Create a new blog

To create that directory you would type:
```shell
$ fuel new dogblog
```
That command tells fuel to create a new installation that's setup for creating your blog.  You'll have the following directories:

~/dogblog/content
~/dogblog/config
~/dogblog/views
~/dogblog/style

## The content directory

The content directory contains all your markdown files.  The directory structure here is mirrored in your site.  For example if you have:

content/
  index.md

Then that index.md will be used to build your index.html.  Let's say you want a very traditional site layout:

content/
  index.md
  about/
    index.md
  dogs/
    cairnterrier.md
    pug.md
  contact/
    index.md

Then you'll get the following web site:

/
  index.html
    about/
      index.html
    dogs/
      cairnterrier.html
      pug.html
    contact/
      index.html
  style/

## The config directory

For future use, not used in the current release.

## The views directory

## The style directory

# Adding new content

# Adding a style

## How to use layouts

## How to use templates

## How to use partial templates

# Using JSON-enhanced markdown with Fuel

## What is JSON

## Example use of JSON-enhanced markdown

# Storing your content in a database

# Troubleshooting

## Verbose mode
