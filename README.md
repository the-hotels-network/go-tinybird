# Go-Tinybird

A [Tinybird](https://www.tinybird.co/) module for Go. It provides an easy and standard way of getting data through the Tinybird API.

## Features

- Lightweight and fast. Native Go implementation, no C-bindings.
- Connection pooling for Tinybird.
- JSON, [NDJSON](http://ndjson.org/) and CSV response formats.
- Parallel HTTP requests.
- Send [events](https://www.tinybird.co/docs/forward/get-data-in/events-api) to data sources.
- Read and write values in nested JSON structures using dot-path access.
- Auto cast string values to their inferred Go types (`int64`, `float64`).
- Custom logic `before` and `after` request execution (e.g. caching).
- Shared logger with [logrus](https://github.com/sirupsen/logrus).
- Test your code with mocks.

## Requirements

Go 1.24 or higher.

## Installation

```bash
go get -u github.com/the-hotels-network/go-tinybird
```

## Quickstart

```go
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

## Working with Data

The `Data` type (`[]Row`) provides methods to read and write values using dot-separated paths:

```go
res := req.Response

// Get a value by key
val := res.Data.FetchOne("revenue")

// Get a nested value using dot-path
val := res.Data.Get("stats.daily.total")

// Set a value
res.Data.Set("stats.daily.total", 100)

// Set with auto cast (converts numeric strings to int64 or float64)
res.Data.SetWithAutoCast("price", "99.99")   // stores float64(99.99)
res.Data.SetWithAutoCast("count", "42")       // stores int64(42)
res.Data.SetWithAutoCast("label", "hello")    // stores "hello" (unchanged)
```

The `AutoCast` function can also be used standalone:

```go
tinybird.AutoCast("42")    // int64(42)
tinybird.AutoCast("3.14")  // float64(3.14)
tinybird.AutoCast("hello") // "hello"
tinybird.AutoCast(99)      // 99 (non-string values pass through)
```

## Parallel Requests

```go
reqs := tinybird.Requests{}

reqs.Add(tinybird.Request{
	Method: http.MethodGet,
	Pipe: tinybird.Pipe{
		Name:      "endpoint_bookings",
		Workspace: tinybird.Workspace{Token: "token-demo"},
	},
})

reqs.Add(tinybird.Request{
	Method: http.MethodGet,
	Pipe: tinybird.Pipe{
		Name:      "endpoint_searches",
		Workspace: tinybird.Workspace{Token: "token-demo"},
	},
})

reqs.Execute()

for _, req := range reqs {
	fmt.Println(req.Response.Data)
}
```

## Events

```go
event := tinybird.Event{
	Method:     http.MethodGet,
	Datasource: "my_datasource",
	Workspace:  tinybird.Workspace{Token: "token-demo"},
}

err := event.Send(map[string]any{"key": "value"})
```

## Response Formats

### NDJSON

Set the format on the pipe or via the `TB_NDJSON` environment variable:

```go
req := tinybird.Request{
	Pipe: tinybird.Pipe{
		Name:   "endpoint",
		Format: tinybird.NDJSON,
	},
}
```

### CSV

```go
req := tinybird.Request{
	Pipe: tinybird.Pipe{
		Name:   "endpoint",
		Format: tinybird.CSV,
	},
}
// Access raw body: req.Response.Body
```

## Before / After Hooks

Add custom logic around request execution, for example a cache layer:

```go
req := tinybird.Request{
	Before: func(r *tinybird.Request) error {
		// Check cache before hitting the API
		return nil
	},
	After: func(r *tinybird.Request) error {
		// Store response in cache
		return nil
	},
}
```

## Configuration

| Variable | Default | Description |
|---|---|---|
| `tinybird.URL_BASE` | `https://api.tinybird.co` | Base API URL |
| `tinybird.Client` | default pool | Custom `HTTPClient` |
| `tinybird.MAX_IDLE_CONNS` | - | Max idle connections |
| `tinybird.MAX_CONNS_PER_HOST` | - | Max connections per host |
| `tinybird.MAX_IDLE_CONNS_PER_HOST` | - | Max idle connections per host |
| `tinybird.CONNS_TIMEOUT` | - | Connection timeout (seconds) |

## Tests

```bash
make tests
```

## More Examples

See the [example directory](https://github.com/the-hotels-network/go-tinybird/tree/main/example) for complete working examples.

## Local Development

Point to a local version of the module:

```bash
go mod edit -replace github.com/the-hotels-network/go-tinybird=$HOME/go/src/github.com/the-hotels-network/go-tinybird
```

Revert:

```bash
go mod edit -dropreplace github.com/the-hotels-network/go-tinybird
go get -u
```
