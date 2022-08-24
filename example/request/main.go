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
	params.Add("start_date", "2022-05-01")
	params.Add("end_date", "2022-05-30")
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
	}

	err := req.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res := req.Response
	fmt.Println("Status:", res.Status)
	fmt.Println("Error:", res.Error)
	fmt.Println("Data:", res.Data)
}
