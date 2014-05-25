package nicovideo

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
)

// Download takes video ID and returns content as io.ReadCloser
func (c *Client) Download(ID string) (io.ReadCloser, error) {
	if c.IsLogin == false {
		return nil, errors.New("login before download")
	}

	// Must visit video page to get some cookies.
	_, err := c.Get(fmt.Sprintf("http://www.nicovideo.jp/watch/%s", ID))
	if err != nil {
		return nil, err
	}

	resp, err := c.Get(
		fmt.Sprintf("http://flapi.nicovideo.jp/api/getflv/%s", ID),
	)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	query, err := url.QueryUnescape(string(body))
	if err != nil {
		return nil, err
	}

	vals, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}

	resp, err = c.Get(vals.Get("url"))
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
