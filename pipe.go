package tinybird

import (
	"fmt"
	"net/url"
)

type Pipe struct {
	Name       string
	URL        string
	Parameters url.Values
	Workspace  Workspace
}

// Build and return the pipe URL.
func (p *Pipe) GetURL() string {
	p.URL = fmt.Sprintf(
		"%s/%s.%s",
		URL(),
		p.Name,
		Format(),
	)

	return p.URL
}
