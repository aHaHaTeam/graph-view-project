package models

type Graph struct {
	id          int
	name        string
	description string
}

func (graph *Graph) GetId() int {
	return graph.id
}
