package gui

type Graph struct {
	// gui.Node with default parameters
	defaultNode Node

	// gui.Edge with default parameters
	defaultEdge Edge
}

func NewGraph(defaultNode Node, defaultEdge Edge) *Graph {
	return &Graph{
		defaultNode: defaultNode,
		defaultEdge: defaultEdge,
	}
}
