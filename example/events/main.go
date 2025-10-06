package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/the-hotels-network/go-tinybird"
)

func main() {
	req := tinybird.Request{
		Method: http.MethodGet,
		Event: tinybird.Event{
			Datasource: "ep_quantum_disparities",
			Workspace: tinybird.Workspace{
				Token: os.Getenv("TB_TOKEN"),
			},
		},
	}

	err := req.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// res := req.Response
	// fmt.Println("Status code:", res.Status)
	// fmt.Println("Elapsed time:", req.Elapsed)
	// fmt.Println("Error:", res.Error)
	// fmt.Println("Data:", res.Data)
	// fmt.Println("URL:", req.URI())
}
