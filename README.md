# Go-Tinybird

A [Tinybird](https://www.tinybird.co/) module for Go. Why need this module? It provides an easy and standard way of getting data through the Tinybird API.

## Features

- Lightweight and fast.
- Native Go implementation. No C-bindings, just pure Go
- Connection pooling for HTTP.
- Test your code with mocks.
- Allow JSON, [NDJSON](http://ndjson.org/) and CSV between tinybird and this module.
- Parallelize HTTP requests.
- Add custom logic `after` and `before` execute request. For example a cache system.
- Shared logger with [logrus](https://github.com/sirupsen/logrus).

## Requirements

Go 1.19 or higher.

## Installation

Simple install the package to your $GOPATH with the go tool from shell:

```bash
go get -u github.com/the-hotels-network/go-tinybird
```

Make sure Git is installed on your machine and in your system's PATH.

## Configure

### NDJSON - Newline-delimited JSON

You can configure it by environment variable `TB_NDJSON` or pipe. By default NDJSON is disabled. Please see this [example to enable via pipe](https://github.com/the-hotels-network/go-tinybird/tree/main/example/request_with_ndjson).

## Quickstart

```
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
	params.Add("property_id", "1234")

	req := tinybird.Request{
		Method: http.MethodGet,
		Pipe: tinybird.Pipe{
			Name:       "tinybird_endpoint",
			Parameters: params,
			Workspace: tinybird.Workspace{
				Token: "token-demo",
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
```

To see more examples, please go to [this directory](https://github.com/the-hotels-network/go-tinybird/tree/main/example).

## Tests

```
make tests
```

## Local usage

To point to the local version of a dependency in Go rather than the one over the web, use the replace keyword.

And now when you compile this module (go install), it will use your local code rather than the other dependency.

```bash
go mod edit -replace github.com/the-hotels-network/go-tinybird=$HOME/go/src/github.com/the-hotels-network/go-tinybird
```

Revert replacement:

```bash
go mod edit -dropreplace github.com/the-hotels-network/go-tinybird
go get -u
```
