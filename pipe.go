package tinybird

import (
	"net/url"
)

type Pipe struct {
	Name       string
	URL        string
	Parameters url.Values
	Workspace  Workspace
}
