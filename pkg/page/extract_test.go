package page

import (
	_ "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/url"
	"os"
	"testing"
)

func TestPage_ExtractURLs(t *testing.T) {
	h, err := os.Open("page1.html")
	require.Nil(t, err)

	root := url.URL{
		Scheme: "http",
		Host:   "www.example.com",
	}

	m := map[string]bool{}
	r := Map(StripDuplicates(m),
		Map(StripFragment,
			Map(SetRoot(root), NewExtractor(h))))
	us, err := r.ReadAll()
	require.Equal(t, io.EOF, err)
	for _, u := range us {
		t.Log(u.String())
	}
}
