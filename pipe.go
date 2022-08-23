package tinybird

import (
	"fmt"
	"net/url"
)

type Pipe struct {
	Name       string
	URL        string
	Parameters url.Values
}

// Build and return the pipe URL.
func (p *Pipe) GetURL() string {
	p.URL = fmt.Sprintf(
		"%s/%s.%s?%s",
		URL(),
		p.Name,
		Format(),
		p.Parameters.Encode(),
	)

	return p.URL
}

// Execute pipe.
func (p *Pipe) Execute() {
	fmt.Println(p.GetURL())
}
