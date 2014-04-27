package foursquare

import (
	"encoding/json"
)

type VenueDetailResponse struct {
	CommonResponse
	Response struct {
		Venue Venue `json:"venue"`
	} `json:"response"`
}

type Venue struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Contact struct {
		Phone          string `json:"phone,omitempty"`
		FormattedPhone string `json:"formattedphone,omitempty"`
	} `json:"contact,omitempty"`
	Location struct {
		Addrees     string  `json:"address,omitempty"`
		CrossStreet string  `json:"crossstreet,omitempty"`
		Lat         float64 `json:"lat,omitempty"`
		Lng         float64 `json:"lng,omitempty"`
		PostalCode  string  `json:"postalCode,omitempty"`
		City        string  `json:"city,omitempty"`
		State       string  `json:"state,omitempty"`
		CountryCode string  `json:"cc,omitempty"`
		Country     string  `json:"country,omitempty"`
	} `json:"location,omitempty"`
	CanonicalURL string     `json:"canonicalUrl,omitempty"`
	Categories   []Category `json:"categories,omitempty"`
	Stats        struct {
		CheckinsCount int `json:"checkinsCount"`
		UsersCount    int `json:"usersCount"`
		TipCount      int `json:"tipCount"`
	} `json:"stats"`
	// TODO(dolapo): tips, etc
}

type Category struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	PluralName string `json:"pluralName"`
	ShortName  string `json:"shortName"`
	Icon       struct {
		Prefix string `json:"prefix"`
		Suffix string `json:"suffix"`
	} `json:"icon"`
}

func (c *Client) VenueDetail(id string) (*VenueDetailResponse, error) {
	resp, err := c.Get("venues/"+id, nil)

	if err != nil {
		return nil, err
	}
	defer resp.Close()

	r := &VenueDetailResponse{}
	json.NewDecoder(resp).Decode(&r)

	return r, nil
}
