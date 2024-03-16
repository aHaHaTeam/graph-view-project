package gui

import (
	"graph-view-project/models"
	"graph-view-project/wasm/physics"
	"image/color"
)

type Edge struct {
	width float32
	color color.Color
	shape models.EdgeShape
}

func (edge *Edge) Draw(begin, end physics.Vec2d) {
	panic("not implemented exception")
}

func (edge *Edge) SetColor(color color.Color) {
	edge.color = color
}

func (edge *Edge) SetShape(shape models.EdgeShape) {
	edge.shape = shape
}
