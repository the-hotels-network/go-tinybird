package tinybird

import (
	"fmt"
)

// Each workspace has a name and token.
type Workspace struct {
	// Workspace name:
	Name string
	// Each workspace have individual token:
	Token string
}

func (w *Workspace) IsSet() bool {
	return len(w.Token) > 0
}

func (w *Workspace) GetToken() string {
	return fmt.Sprintf("Bearer %s", w.Token)
}
