package page

import (
	"golang.org/x/net/html"
	"io"
	url "net/url"
)

type extractor struct {
	t *html.Tokenizer
}

var (
	attrs = map[string]string{
		"a":    "href",
		"img":  "src",
		"form": "action",
		"link": "href",
	}
)

func NewExtractor(data io.Reader) Reader {
	return &extractor{
		t: html.NewTokenizer(data),
	}
}

func (r *extractor) Read() (*url.URL, error) {
	for {
		t := r.t.Next()
		switch t {
		case html.ErrorToken:
			return nil, r.t.Err()
		case html.StartTagToken:
			fallthrough
		case html.SelfClosingTagToken:
			return tokenURL(r.t.Token())
		}
	}
	return nil, nil
}

func (r *extractor) ReadAll() ([]*url.URL, error) {
	var us []*url.URL
	for {
		u, err := r.Read()
		if err != nil {
			return us, err
		}
		if u != nil {
			us = append(us, u)
		}
	}
	return us, nil
}

func tokenURL(t html.Token) (*url.URL, error) {
	an, ok := attrs[t.Data]
	if !ok {
		return nil, nil
	}
	a := findAttr(t, an)
	if a == nil {
		return nil, nil
	}
	return url.Parse(a.Val)
}

func findAttr(t html.Token, n string) *html.Attribute {
	for _, a := range t.Attr {
		if a.Key == n {
			return &a
		}
	}
	return nil
}
