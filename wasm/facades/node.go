package facades

import (
	"graph-view-project/wasm/gui"
	"graph-view-project/wasm/models"
	"graph-view-project/wasm/physics"
	"image/color"
)

type Node struct {
	model models.Node
	point physics.Node
	node  gui.Node
}

func (*Node) Draw() {
	panic("not implemented exception")
}

func (node *Node) Update(c chan struct{}) {
	node.point.Update(c)
}

func (node *Node) Move(time float64) {
	node.point.Move(time)
}

func (node *Node) SetNodeId(id int) {
	node.model.SetId(id)
}

func (node *Node) SetNodeName(name string) {
	node.model.SetName(name)
}

func (node *Node) SetNodeData(data []byte) {
	node.model.SetData(data)
}

func (node *Node) SetNodeColor(color color.Color) {
	node.node.SetColor(color)
}

func (node *Node) SetNodeShape(shape gui.NodeShape) {
	node.node.SetShape(shape)
}

func (node *Node) SetNodeSize(size float32) {
	node.node.SetSize(size)
}
