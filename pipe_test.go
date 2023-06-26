package tinybird_test

import (
	"net/url"
	"testing"

	"github.com/the-hotels-network/go-tinybird"

	"github.com/stretchr/testify/assert"
)

func TestGetParameters(t *testing.T) {
	params := url.Values{}
	params.Add("start_date", "2022-05-01")
	params.Add("end_date", "2022-05-30")
	params.Add("id", "1")
	params.Add("id", "2")
	params.Add("id", "3")
	params.Add("id", "4")

	p := tinybird.Pipe{
		Name:       "ep_test",
		Alias:      "test",
		Parameters: params,
		Workspace: tinybird.Workspace{
			Name: "test",
		},
	}

	assert.Equal(t, p.GetParameters(), "end_date=2022-05-30&id=1%2C2%2C3%2C4&start_date=2022-05-01")
}
