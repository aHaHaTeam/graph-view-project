package facades

import (
	"graph-view-project/models"
	"graph-view-project/wasm/gui"
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

func (node *Node) Update(c chan struct{}, nodes *[]physics.Node, graph *physics.Graph) {
	node.point.Update(c, nodes, graph)
}

func (node *Node) Move(time float64) {
	node.point.Move(time)
}

func (node *Node) SetNodeId(id int) {
	node.model.Id = id
}

func (node *Node) SetNodeName(name string) {
	node.model.Name = name
}

func (node *Node) SetNodeData(data []byte) {
	node.model.Data = data
}

func (node *Node) SetNodeColor(color color.Color) {
	node.node.SetColor(color)
}

func (node *Node) SetNodeShape(shape models.NodeShape) {
	node.node.SetShape(shape)
}

func (node *Node) SetNodeSize(size float32) {
	node.node.SetSize(size)
}
