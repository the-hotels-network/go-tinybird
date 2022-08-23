package tinybird

type Pipes []Pipe

// Add new pipe.
func (p *Pipes) Add(n Pipe) {
	if !p.Has(n) {
		(*p) = append((*p), n)
	}
}

// Check pipe if exist.
func (p Pipes) Has(n Pipe) bool {
	for _, i := range p {
		if i.Name == n.Name {
			return true
		}
	}

	return false
}
