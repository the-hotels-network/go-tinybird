package tinybird

import (
	"fmt"
)

type Event struct {
	Name       string
	Datasource string
	Workspace  Workspace
}

func (e Event) GetURL() string {
	return URL_BASE
}

func (e Event) GetURI() string {
	return fmt.Sprintf(
		"%s/v0/events?name=%s",
		e.GetURL(),
		e.Datasource,
	)
}
