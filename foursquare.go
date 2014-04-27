package foursquare

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// TODO(dolapo): Consider making this configurable
const (
	BASE_API_URL   = "api.foursquare.com"
	CLIENT_VERSION = "20140417"
)

type Client struct {
	accessToken  string
	clientId     string
	clientSecret string
}

func NewClient(accessToken string) *Client {
	return &Client{
		accessToken: accessToken,
	}
}

func NewUserlessClient(clientId string, clientSecret string) *Client {
	return &Client{
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (c *Client) resource(path string, params interface{}) url.URL {
	bp := map[string][]string{
		"v": []string{CLIENT_VERSION},
	}
	if len(c.accessToken) > 0 {
		bp["oauth_token"] = []string{c.accessToken}
	} else {
		bp["client_id"] = []string{c.clientId}
		bp["client_secret"] = []string{c.clientSecret}
	}

	// TODO(dolapo): location params..

	u := url.URL{
		Scheme:   "https",
		Host:     BASE_API_URL,
		Path:     "/v2/" + path,
		RawQuery: url.Values(bp).Encode(),
	}

	if params != nil {
		q, _ := query.Values(params)
		u.RawQuery += "&" + q.Encode()
	}

	return u
}

func (c *Client) Get(path string, params interface{}) (io.ReadCloser, error) {
	u := c.resource(path, params)
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		slurp, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, slurp)
	}

	return resp.Body, nil
}

type Notification struct {
	Type string `json:"type"`
	Item struct {
		UnreadCount int `json:"unreadCount"`
	} `json:"item,omitempty"`
}

type CommonResponse struct {
	Meta struct {
		Code int `json:"code"`
	} `json:"meta"`
	Notifications []Notification `json:notifications,omitempty"`
}
