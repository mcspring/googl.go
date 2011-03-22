# googl.go

googl.go is an open source implementation of Google URL Shortener API service writen in go

googl.go is licensed under the Apache Lincense, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).

## Installation

1. make sure you have a working go environment ( See the [Install Instructions](http://golang.org/doc/install.html) ).
2. `git clone git://github.com/mcspring/googl.go.git`
3. `cd googl.go` && `make install`

## Example

	import (
		"os"
		"fmt"
		"googl"
	)

	const (
		API_GOOGL_KEY = "" // your google url shorten api key, default is empty
	)

	func main() {
		longUrl := "http://www.google.com"

		urlapi := googl.NewGoogl(GOOGL_KEY)

		shortUrl, err := urlapi.Shorten(longUrl)
		if nil != err {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Short URL: %s\n", shortUrl)

		longUrl, err = urlapi.Expand(shortUrl)
		if nil != err {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Long URL: %s\n", longUrl)
	}

This will output:

	Short URL: http://goo.gl/fbsS
	Long URL: http://www.google.com/

## Attention

googl.go only implement [url.insert](http://code.google.com/apis/urlshortener/v1/reference.html#method_urlshortener_url_insert) 
and [url.get](http://code.google.com/apis/urlshortener/v1/reference.html#method_urlshortener_url_get) of Google URL Shortener API now.

## About

googl.go is written by [Spring Mc](https://github.com/mcspring).

And you can follow me on [Twitter](http://twitter.com/mcspring)!

