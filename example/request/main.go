package main

import (
	"net/url"

	"github.com/the-hotels-network/go-tinybird"
)

func main() {
	ws := tinybird.Workspace{
		Name: "demo",
		Token: "",
	}

	params := url.Values{}
	params.Add("start_date", "2022-01-01")
	params.Add("end_date", "2022-01-30")

	ws.Pipes.Add(
		tinybird.Pipe{
			Name:"foo",
			Parameters: params,
		},
	)

	pipe := ws.Pipe("foo")
	pipe.Execute()
}
