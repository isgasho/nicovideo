package nicovideo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/url"
)

// RankingGenre is type of genre.
type RankingGenre string

// Ranking genres.
const (
	RankingGenreAll RankingGenre = "all"
)

// RankingType is type of ranking type.
type RankingType string

// Ranking types.
const (
	RankingTypeView RankingType = "view"
)

// RankingSpan is type of ranking span.
type RankingSpan string

// Ranking spans.
const (
	RankingSpanDaily RankingSpan = "daily"
)

// Ranking represents ranking.
type Ranking struct {
	Status     string `xml:"status,attr"`
	Count      int    `xml:"count"`
	VideoInfos []struct {
		Video struct {
			ID        string `xml:"id"`
			IsDeleted bool   `xml:"deleted"`
		} `xml:"video"`
	} `xml:"video_info"`
}

// GetRanking returns array of video IDs.
func (c *Client) GetRanking(
	typ RankingType,
	genre RankingGenre,
	span RankingSpan,
) (*Ranking, error) {
	apiURL, _ := url.Parse(
		"http://api.ce.nicovideo.jp/nicoapi/v1/video.ranking",
	)

	params := url.Values{
		"type":  {string(typ)},
		"genre": {string(genre)},
		"span":  {string(span)},
	}

	apiURL.RawQuery = params.Encode()

	resp, err := c.Get(apiURL.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ranking := &Ranking{}
	err = xml.Unmarshal(body, &ranking)
	if err != nil {
		return nil, err
	}

	if ranking.Status != "ok" {
		return nil, fmt.Errorf("api status: %s", ranking.Status)
	}

	return ranking, nil
}

// GetDailyAllRanking returns array of video IDs.
func (c *Client) GetDailyAllRanking() (*Ranking, error) {
	return c.GetRanking(
		RankingTypeView,
		RankingGenreAll,
		RankingSpanDaily,
	)
}
