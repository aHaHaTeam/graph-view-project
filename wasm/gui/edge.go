package gui

import "image/color"

type EdgeShape int

const (
	Arrow EdgeShape = iota
	Line
)

type Edge struct {
	color color.Color
	shape EdgeShape
}

func (e *Edge) Draw() {
	panic("not implemented exception")
}

func (e *Edge) SetColor(color color.Color) {
	e.color = color
}

func (e *Edge) SetShape(shape EdgeShape) {
	e.shape = shape
}
