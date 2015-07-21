package google_url_shortener

import (
	"testing"
)

func TestPostUrl(t *testing.T) {
	g := new(GoogleUrlShortener)
	if g.postUrl() == "" {
		panic("Could not get post url")
	}
}

func TestExpandedUrlWithHttp(t *testing.T) {
	g := new(GoogleUrlShortener)
	expanded_url, err := g.ExpandedUrl("http://goo.gl/Ri16S");
	if expanded_url != "https://sites.google.com/site/nycjv321/" {
		panic(err)
	}
}

func TestExpandedUrlWithHttps(t *testing.T) {
	g := new(GoogleUrlShortener)
	expanded_url, err := g.ExpandedUrl("https://goo.gl/Ri16S");
	if expanded_url != "https://sites.google.com/site/nycjv321/" {
		panic(err)
	}
}


func TestInvalidExpandedUrl(t *testing.T) {
	g := new(GoogleUrlShortener)
	expanded_url, err := g.ExpandedUrl("blah");
	if err == nil {
		panic("Expected error when trying to expand invalid url")
	}
	_ = expanded_url
}

func TestShortenedUrl(t *testing.T) {
	g := new(GoogleUrlShortener)
	shortened_url, err := g.ShortenedUrl("https://sites.google.com/site/nycjv321/");
	if shortened_url != "https://goo.gl/lUXm11" {
		panic(err)
	}
}

func TestInvalidShortenedUrl(t *testing.T) {
	g := new(GoogleUrlShortener)
	shortened_url, err := g.ShortenedUrl("blah/");
	if shortened_url != "http://goo.gl/XJx8gw" {
		panic(err)
	}
}

