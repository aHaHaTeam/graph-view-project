package gui

import (
	"graph-view-project/models"
	"image/color"
)

var DefaultNode = Node{
	size: 1.0,
	color: color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	},
	shape: models.NodeShape(0),
}

var DefaultEdge = Edge{
	width: 0.1,
	color: color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	},
	shape: models.EdgeShape(0),
}

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
