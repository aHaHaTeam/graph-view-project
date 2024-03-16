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

func NewEdge(width float32, color color.Color, shape models.EdgeShape) *Edge {
	return &Edge{width: width, color: color, shape: shape}
}

func (edge *Edge) Draw(begin, end physics.Vec2d) {
	panic("not implemented exception")
}

func (edge *Edge) Width() float32 {
	return edge.width
}

func (edge *Edge) SetWidth(width float32) {
	edge.width = width
}

func (edge *Edge) Color() color.Color {
	return edge.color
}

func (edge *Edge) SetColor(color color.Color) {
	edge.color = color
}
func (edge *Edge) Shape() models.EdgeShape {
	return edge.shape
}

func (edge *Edge) SetShape(shape models.EdgeShape) {
	edge.shape = shape
}
