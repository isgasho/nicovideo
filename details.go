package nicovideo

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
)

// VideoDetails contains details from API.
type VideoDetails struct {
	Title       string `xml:"thumb>title"`
	Description string `xml:"thumb>description"`
	Type        string `xml:"thumb>movie_type"`
	Error       string `xml:"error>code"`
}

// GetVideoDetails takes video ID and returns details.
func (c *Client) GetVideoDetails(ID string) (*VideoDetails, error) {
	resp, err := c.Get(
		fmt.Sprintf("http://ext.nicovideo.jp/api/getthumbinfo/%s", ID),
	)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	details := &VideoDetails{}
	err = xml.Unmarshal(body, &details)
	if err != nil {
		return nil, err
	}

	if details.Error != "" {
		return nil, errors.New(details.Error)
	}

	return details, nil
}
