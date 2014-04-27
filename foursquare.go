package foursquare

import (
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// TODO(dolapo): Consider making this configurable
const (
	BASE_API_URL   = "https://api.foursquare.com/"
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
	u := url.URL{
		Scheme: "https",
		Host:   BASE_API_URL,
		Path:   "/v1/" + path,
	}

	if params != nil {
		q, _ := query.Values(params)
		u.RawQuery += "&" + q.Encode() + "&"
	}
	u.RawQuery += url.Values(map[string][]string{
		"v":           []string{CLIENT_VERSION},
		"oauth_token": []string{c.accessToken},
	}).Encode()

	return u
}

func (c *Client) doGet(path string, params interface{}) (io.ReadCloser, error) {
	u := c.resource(path, params)
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

/*
func (c *Client) VenuesSearch(params *VenuesSearchParams) *VenuesSearchResult {
} */
