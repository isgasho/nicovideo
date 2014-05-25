package nicovideo

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
)

// Login takes email and password and returns error if fail.
func (c *Client) Login(email, password string) error {
	if c.IsLogin {
		return nil
	}

	params := url.Values{
		"mail":     {email},
		"password": {password},
	}

	req, _ := http.NewRequest(
		"POST",
		"https://secure.nicovideo.jp/secure/login",
		bytes.NewBufferString(params.Encode()),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Go's redirection handling just sucks...
	_, err := c.Do(req)
	if err == nil {
		return errors.New("invalid email/password")
	}

	if httpError, ok := err.(*url.Error); ok {
		if httpError.Err.Error() == redirectError {
			c.IsLogin = true
			return nil
		}
	}

	return err
}
