package facades

import (
	"graph-view-project/models"
	"graph-view-project/wasm/content"
	"graph-view-project/wasm/gui"
	"graph-view-project/wasm/physics"
	"image/color"
)

type Node struct {
	id      int
	edges   []*Edge
	content content.Node
	point   physics.Node
	node    gui.Node
}

func NewNode(model *models.Node) *Node {
	return &Node{
		id: model.Id,

		content: *content.NewNode(
			model.Name,
			model.Data,
		),
		point: *physics.NewNode(),
		node:  gui.Node{},
	}
}

func (n *Node) GetModel() models.Node {
	return models.Node{
		Id:    n.id,
		Name:  n.content.Name(),
		Data:  n.content.Data(),
		Size:  n.node.Size(),
		Color: n.node.Color(),
		Shape: n.node.Shape(),
	}
}

func (*Node) Draw() {
	panic("not implemented exception")
}

func (n *Node) Update(c chan struct{}, nodes *[]physics.Node, graph *physics.Graph) {
	n.point.Update(c, nodes, graph)
}

func (n *Node) Move(time float64) {
	n.point.Move(time)
}

func (n *Node) SetNodeId(id int) {
	n.id = id
}

func (n *Node) SetNodeName(name string) {
	n.content.SetName(name)
}

func (n *Node) SetNodeData(data []byte) {
	n.content.SetData(data)
}

func (n *Node) SetNodeColor(color color.Color) {
	n.node.SetColor(color)
}

func (n *Node) SetNodeShape(shape models.NodeShape) {
	n.node.SetShape(shape)
}

func (n *Node) SetNodeSize(size float32) {
	n.node.SetSize(size)
}
