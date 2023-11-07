package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/the-hotels-network/go-tinybird"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	// Configure tinybird logger to integrate with Logrus.
	tinybird.NewLogger(tinybird.LogrusAdapter{})

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
		Format: tinybird.CSV,
	}

	err := req.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res := req.Response
	fmt.Println("Status code:", res.Status)
	fmt.Println("Elapsed time:", req.Elapsed)
	fmt.Println("Error:", res.Error)
	fmt.Println("Data:", res.Body)
}
