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
	size  float32
	color color.Color
	shape NodeShape
}

func (n *Node) Draw() {
	panic("not implemented exception")
}

func (n *Node) SetColor(color color.Color) {
	n.color = color
}

func (n *Node) SetShape(shape NodeShape) {
	n.shape = shape
}
