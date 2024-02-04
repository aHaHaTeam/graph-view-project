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

func (e *Edge) Draw(begin, end physics.Vec2d) {
	panic("not implemented exception")
}

func (e *Edge) SetColor(color color.Color) {
	e.color = color
}

func (e *Edge) SetShape(shape EdgeShape) {
	e.shape = shape
}
