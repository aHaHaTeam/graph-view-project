package models

import "image/color"

type Graph struct {
	Id          int
	Name        string
	Description string
	Nodes       []*Node
	Edges       []*Edge

	// Default gui.Node parameters
	DefaultNodeSize  float32
	DefaultNodeColor color.Color
	DefaultNodeShape NodeShape

	// Default gui.Edge parameters
	DefaultEdgeWidth float32
	DefaultEdgeColor color.Color
	DefaultEdgeShape EdgeShape
}

func NewGraph(id int,
	name string,
	description string,
	nodes []*Node,
	edges []*Edge,
	defaultNodeSize float32,
	defaultNodeColor color.Color,
	defaultNodeShape NodeShape,
	defaultEdgeWidth float32,
	defaultEdgeColor color.Color,
	defaultEdgeShape EdgeShape,
) *Graph {
	return &Graph{
		Id:          id,
		Name:        name,
		Description: description,
		Nodes:       nodes, Edges: edges,
		DefaultNodeSize:  defaultNodeSize,
		DefaultNodeColor: defaultNodeColor,
		DefaultNodeShape: defaultNodeShape,
		DefaultEdgeWidth: defaultEdgeWidth,
		DefaultEdgeColor: defaultEdgeColor,
		DefaultEdgeShape: defaultEdgeShape,
	}
}
