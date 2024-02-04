package facades

import (
	"graph-view-project/wasm/gui"
	"graph-view-project/wasm/models"
	"graph-view-project/wasm/physics"
	"image/color"
)

type Edge struct {
	model   models.Edge
	segment physics.Edge
	edge    gui.Edge
}

func (*Edge) Draw() {
	panic("not implemented exception")
}

func (*Edge) SetEdgeName(name string) {
	panic("not implemented exception")
}

func (*Edge) SetNodeDescription(description string) {
	panic("not implemented exception")
}

func (*Edge) SetNodeColor(color color.Color) {
	panic("not implemented exception")
}

func (*Edge) SetNodeShape(data gui.EdgeShape) {
	panic("not implemented exception")
}
