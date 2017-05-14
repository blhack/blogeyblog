# BLOGEYBLOG is my blogging software.

## how to use it:

```
cd blogeyBlog
go get -d ./...
go build -o blogeyBlog blogeyBlog.go
./blogeyBlog
```

## Formatting blog posts

```
touch ./blogs/test.blogeyBlog
```

Then format the file, ./blogs/test.blogeyBlog as follows:

```
This is the title of the blog

This is the content of the blog.

I gave two newlines there, so this is a new paragraph in the rendered HTML document.
This only had one newline, so it is still the same paragraph.

You can make a [link](http://blhack.me/blhack) by placing the link text between square brackets, and the link destination immediately after it in parenthesis.

![This is some alt text](http://example.com/images/anImage.jpg)

Images can be linked in a similar way to hrefs, if they are prepended by the ! character.

Good luck.
