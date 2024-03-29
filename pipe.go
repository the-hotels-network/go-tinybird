package tinybird

import (
	"net/url"
	"strings"
)

// Pipe is a object on tinybird, contains one or more SQL queries (Nodes) that
// result in either an API endpoint or a Materialized View.
type Pipe struct {
	// Name of pipe. Define by user.
	Name string
	// Alias of pipe. Optional.
	Alias string
	// Is for internal purpose and is generated by code.
	URL string
	// Parameters use on pipe. Define by user.
	Parameters url.Values
	// Is an area that contains a set of Tinybird resources, including Pipes,
	// Nodes, APIs, Data Sources & Auth Tokens. Define by user.
	Workspace Workspace
}

func (p *Pipe) GetParameters() string {
	for key, value := range p.Parameters {
		if len(value) > 1 {
			p.Parameters.Del(key)
			p.Parameters.Add(key, strings.Join(value, ","))
		}
	}

	return strings.Replace(p.Parameters.Encode(), "+", "%20", -1)
}
