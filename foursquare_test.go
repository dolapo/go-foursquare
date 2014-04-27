package foursquare

import (
	"flag"
	"io/ioutil"
	"testing"
)

var (
	accessToken = flag.String("access_token", "", "Access Token")
)

func TestGet(t *testing.T) {
	flag.Parse()

	client := NewClient(*accessToken)
	resp, err := client.Get("venues/4ef0e7cf7beb5932d5bdeb4e", nil)

	if err != nil {
		t.Errorf("error: %v", err)
		return
	}
	defer resp.Close()

	if b, err := ioutil.ReadAll(resp); err == nil {
		t.Logf("%+v", string(b))
	}
}

func TestVenueDetail(t *testing.T) {
	flag.Parse()
	client := NewClient(*accessToken)
	venue, err := client.VenueDetail("4ef0e7cf7beb5932d5bdeb4e")

	if err != nil {
		t.Errorf("error: %v", err)
		return
	}
	t.Logf("%+v", venue)
}
