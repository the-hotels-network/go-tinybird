package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

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

	ws := tinybird.Workspace{
		Token: os.Getenv("TB_TOKEN"),
	}

	reqs := tinybird.Requests{}
	req1 := tinybird.Request{
		Method: http.MethodGet,
		Pipe: tinybird.Pipe{
			Name:       "ep_quantum_bookings",
			Parameters: params,
			Workspace:  ws,
		},
	}
	req2 := tinybird.Request{
		Method: http.MethodGet,
		Pipe: tinybird.Pipe{
			Name:       "ep_quantum_searches",
			Parameters: params,
			Workspace:  ws,
		},
	}

	reqs.Add(req1)
	reqs.Add(req2)
	reqs.Execute()

	fmt.Println("Total elapsed time:", reqs.Duration())
	for _, req := range reqs {
		fmt.Println(strings.Repeat("-", 40))
		fmt.Println("Pipe name:", req.Pipe.Name)
		fmt.Println("Elapsed time:", req.Elapsed)
		fmt.Println("Status:", req.Response.Status)
		fmt.Println("Error:", req.Response.Error)
		fmt.Println("Data:", req.Response.Data)
	}
}
