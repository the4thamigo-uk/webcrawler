package page

import (
	url "net/url"
)

type Reader interface {
	Read() (*url.URL, error)
	ReadAll() ([]*url.URL, error)
}
