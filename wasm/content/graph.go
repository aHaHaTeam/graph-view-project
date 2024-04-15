package content

type Graph struct {
	name        string
	description string
}

func NewGraph(name string, description string) *Graph {
	return &Graph{name: name, description: description}
}

func (g *Graph) Name() string {
	return g.name
}

func (g *Graph) SetName(name string) {
	g.name = name
}

func (g *Graph) Description() string {
	return g.description
}

func (g *Graph) SetDescription(description string) {
	g.description = description
}
