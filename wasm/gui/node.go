package gui

import (
	"graph-view-project/models"
	"image/color"
)

type Node struct {
	color color.Color
	shape models.NodeShape
	size  float32
}

func (node *Node) Draw() {
	panic("not implemented exception")
}

func (node *Node) SetColor(color color.Color) {
	node.color = color
}

func (node *Node) SetShape(shape models.NodeShape) {
	node.shape = shape
}

func (node *Node) SetSize(size float32) {
	node.size = size
}
