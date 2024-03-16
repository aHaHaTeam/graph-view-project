package gui

import (
	"graph-view-project/models"
	"image/color"
)

type Node struct {
	size  float32
	color color.Color
	shape models.NodeShape
}

func NewNode(size float32, color color.Color, shape models.NodeShape) *Node {
	return &Node{size: size, color: color, shape: shape}
}

func (node *Node) Draw() {
	panic("not implemented exception")
}

func (node *Node) Size() float32 {
	return node.size
}

func (node *Node) SetSize(size float32) {
	node.size = size
}

func (node *Node) Color() color.Color {
	return node.color
}

func (node *Node) SetColor(color color.Color) {
	node.color = color
}

func (node *Node) Shape() models.NodeShape {
	return node.shape
}

func (node *Node) SetShape(shape models.NodeShape) {
	node.shape = shape
}
