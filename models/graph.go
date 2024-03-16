package models

import "image/color"

type Graph struct {
	Id          int
	Name        string
	Description string
	Nodes       []Node
	Edges       []Edge

	// Default gui.Node parameters
	DefaultNodeSize  float32
	DefaultNodeColor color.Color
	DefaultNodeShape NodeShape

	// Default gui.Edge parameters
	DefaultEdgeWidth float32
	DefaultEdgeColor color.Color
	DefaultEdgeShape EdgeShape
}
