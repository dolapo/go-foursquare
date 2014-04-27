package foursquare

import (
	"encoding/json"
)

type UserDetailResponse struct {
	CommonResponse
	Response struct {
		User User `json:"user"`
	} `json:"response"`
}

type User struct {
	Id           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Gender       string `json:"gender"`
	Relationship string `json:"relationship,omitempty"`
	Photo        struct {
		Prefix string `json:"prefix"`
		Suffix string `json:"suffix"`
	} `json:"photo"`
}

func (c *Client) UserDetail(id string) (*UserDetailResponse, error) {
	resp, err := c.Get("users/"+id, nil)

	if err != nil {
		return nil, err
	}
	defer resp.Close()

	r := &UserDetailResponse{}
	json.NewDecoder(resp).Decode(&r)

	return r, nil
}
