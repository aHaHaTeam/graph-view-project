package models

type Graph struct {
	id          int
	name        string
	description string
	Nodes       []Node
	Edges       []Edge
}

func (graph *Graph) GetId() int {
	return graph.id
}
