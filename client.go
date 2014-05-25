package nicovideo

import (
	"code.google.com/p/go.net/publicsuffix"
	"errors"
	"net/http"
	"net/http/cookiejar"
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
			publicsuffix.List,
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
