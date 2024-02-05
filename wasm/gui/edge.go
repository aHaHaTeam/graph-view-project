package gui

import (
	"graph-view-project/wasm/physics"
	"image/color"
)

type EdgeShape int

const (
	Arrow EdgeShape = iota
	Line
)

type Edge struct {
	width float32
	color color.Color
	shape EdgeShape
}

func (edge *Edge) Draw(begin, end physics.Vec2d) {
	panic("not implemented exception")
}

func (edge *Edge) SetColor(color color.Color) {
	edge.color = color
}

func (edge *Edge) SetShape(shape EdgeShape) {
	edge.shape = shape
}
