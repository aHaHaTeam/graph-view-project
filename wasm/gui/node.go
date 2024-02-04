package gui

import (
	"image/color"
)

type NodeShape int

const (
	Circle NodeShape = iota
	Square
)

type Node struct {
	color color.Color
	shape NodeShape
	size  float32
}

func (node *Node) Draw() {
	panic("not implemented exception")
}

func (node *Node) SetColor(color color.Color) {
	node.color = color
}

func (node *Node) SetShape(shape NodeShape) {
	node.shape = shape
}

func (node *Node) SetSize(size float32) {
	node.size = size
}
