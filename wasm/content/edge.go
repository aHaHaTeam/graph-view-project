package content

type Edge struct {
	name        string
	description string
}

func NewEdge(name string, description string) *Edge {
	return &Edge{name: name, description: description}
}

func (e *Edge) Name() string {
	return e.name
}

func (e *Edge) SetName(name string) {
	e.name = name
}

func (e *Edge) Description() string {
	return e.description
}

func (e *Edge) SetDescription(description string) {
	e.description = description
}
