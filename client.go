package nicovideo

import (
	"errors"
	"net/http"
	"net/http/cookiejar"

	"code.google.com/p/go.net/publicsuffix"
)

// Client manages cookies.
type Client struct {
	*http.Client
	IsLogin bool
}

const redirectError = "no redirect"

// NewClient returns a pointer to a new Client instance.
func NewClient() *Client {
	jar, _ := cookiejar.New(
		&cookiejar.Options{
			PublicSuffixList: publicsuffix.List,
		},
	)

	client := &http.Client{
		Jar: jar,
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request,
		) error {
			return errors.New(redirectError)
		},
	}

	return &Client{Client: client}
}
