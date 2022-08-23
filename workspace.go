package tinybird

// Each workspace has a name and token.
type Workspace struct {
	// Workspace name:
	Name string
	// Each workspace have individual token:
	Token string
	// List of pipes inside in workspace:
	Pipes Pipes
}

// Select the pipe to use.
func (w *Workspace) Pipe(name string) *Pipe {
	for i := range w.Pipes {
		if w.Pipes[i].Name == name {
			return &w.Pipes[i]
		}
	}

	return nil
}
