package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/the-hotels-network/go-tinybird"
)

func main() {
	params := url.Values{}
	params.Add("start_date", "2023-05-01")
	params.Add("end_date", "2023-05-30")
	params.Add("property_id", "1011163")
	params.Add("property_id", "1011832")
	params.Add("property_id", "1011846")
	params.Add("property_id", "1011847")

	req := tinybird.Request{
		Method: http.MethodGet,
		Pipe: tinybird.Pipe{
			Name:       "ep_quantum_disparities",
			Parameters: params,
			Workspace: tinybird.Workspace{
				Name:  "quantum",
				Token: os.Getenv("TB_TOKEN"),
			},
		},
		Before: func(r *tinybird.Request) bool {
			// For example, you can check request is saved on redis here.
			fmt.Println("==> Before execute request")
			fmt.Println(" - URL: ", r.URI())

			// Use return true to execute request, and false to not.
			return true
		},
		After: func(r *tinybird.Request) {
			// For example, you can save body request on redis here.
			fmt.Println("==> After execute request")
			fmt.Println(" - Status code:", r.Response.Status)
			fmt.Println(" - Error:", r.Response.Error)
			fmt.Println(" - Data:", r.Response.Data)
		},
	}

	fmt.Println("==> Execute request")
	err := req.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("==> End")
	fmt.Println(" - Elapsed time:", req.Elapsed)
}
