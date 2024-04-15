package models

import "image/color"

type Graph struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Nodes       []*Node `json:"nodes"`
	Edges       []*Edge ` json:"edges"`

	// Default gui.Node parameters
	DefaultNodeSize  float32    `json:"defaultNodeSize"`
	DefaultNodeColor color.RGBA `json:"defaultNodeColor"`
	DefaultNodeShape NodeShape  `json:"defaultNodeShape"`

	// Default gui.Edge parameters
	DefaultEdgeWidth float32    `json:"defaultEdgeWidth"`
	DefaultEdgeColor color.RGBA `json:"defaultEdgeColor"`
	DefaultEdgeShape EdgeShape  `json:"defaultEdgeShape"`
}

func NewGraph(id int,
	name string,
	description string,
	nodes []*Node,
	edges []*Edge,
	defaultNodeSize float32,
	defaultNodeColor color.RGBA,
	defaultNodeShape NodeShape,
	defaultEdgeWidth float32,
	defaultEdgeColor color.RGBA,
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
