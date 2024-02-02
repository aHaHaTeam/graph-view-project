package facades

import (
	"graph-view-project/wasm/gui"
	"graph-view-project/wasm/models"
	"graph-view-project/wasm/physics"
	"image/color"
)

type Node struct {
	model models.Node
	point physics.Point
	node  gui.Node
}

func (*Node) Draw() {
	panic("not implemented exception")
}

func (*Node) Update() {
	panic("not implemented exception")
}

func (*Node) SetNodeName(name string) {
	panic("not implemented exception")
}

func (*Node) SetNodeData(data string) {
	panic("not implemented exception")
}

func (*Node) SetNodeColor(color color.Color) {
	panic("not implemented exception")
}

func (*Node) SetNodeShape(data gui.NodeShape) {
	panic("not implemented exception")
}
