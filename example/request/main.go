package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/the-hotels-network/go-tinybird"
)

func main() {
	params := url.Values{}
	params.Add("start_date", "2023-05-01")
	params.Add("end_date", "2023-05-29")
	params.Add("user_id", "1011")
	params.Add("user_id", "1012")

	req := tinybird.Request{
		Method: http.MethodGet,
		Pipe: tinybird.Pipe{
			Name:       "ep_users",
			Parameters: params,
			Workspace: tinybird.Workspace{
				Name:  "example",
				Token: "testoken",
			},
		},
	}

	tinybird.MockResponseFunc(
		http.StatusOK,
		func(url string) string {
			return `{"id": "01H3EHWCP9TF5GSHYBVFXSZX9M", "type": "user", "min": 0, "currency": "EUR"}`
		},
		nil,
	)

	req.Execute()
	res := req.Response

	fmt.Println("Status code:", res.Status)
	fmt.Println("Elapsed time:", req.Elapsed)
	fmt.Println("Error:", res.Error)
	fmt.Println("Body:", res.Body)
	fmt.Println("URL:", req.URL())
	fmt.Println("URI:", req.URI())
}
