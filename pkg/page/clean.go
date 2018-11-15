package page

import (
	"net/url"
)

type mapper struct {
	r Reader
	f func(u *url.URL) (*url.URL, error)
}

func Map(f func(u *url.URL) (*url.URL, error), r Reader) Reader {
	return &mapper{
		r: r,
		f: f,
	}
}

func (m *mapper) Read() (*url.URL, error) {
	u, err := m.r.Read()
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, nil
	}
	return m.f(u)
}
func (m *mapper) ReadAll() ([]*url.URL, error) {
	var us []*url.URL
	for {
		u, err := m.Read()
		if err != nil {
			return us, err
		}
		if u != nil {
			us = append(us, u)
		}
	}
	return us, nil
}

func SetRoot(root url.URL) func(*url.URL) (*url.URL, error) {
	return func(rel *url.URL) (*url.URL, error) {
		if rel.IsAbs() {
			return rel, nil
		}
		return root.ResolveReference(rel), nil
	}
}

func StripFragment(u *url.URL) (*url.URL, error) {
	u.Fragment = ""
	return u, nil
}

// better to use redis for this due to
// map just getting bigger an bigger
func StripDuplicates(m map[string]bool) func(*url.URL) (*url.URL, error) {
	return func(u *url.URL) (*url.URL, error) {
		s := u.String()
		_, ok := m[s]
		if ok {
			return nil, nil
		}
		m[s] = true
		return u, nil
	}
}
