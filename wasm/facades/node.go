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
	content *content.Node
	point   *physics.Node
	node    *gui.Node
}

func newNode(model *models.Node) *Node {
	return &Node{
		id: model.Id,
		content: content.NewNode(
			model.Name,
			model.Data,
		),
		point: physics.NewNode(),
		node:  gui.NewNode(model.Size, model.Color, model.Shape),
	}
}

func (node *Node) addEdges(adjacentNodes []*Node) {
	adjacentPoints := make([]*physics.Node, len(adjacentNodes))
	for i, adjacentNode := range adjacentNodes {
		if adjacentNode == node {
			adjacentPoints[i] = adjacentNode.point
		} else {
			adjacentPoints[i] = adjacentNode.point
		}
	}
	node.point.AddAdjacentNodes(adjacentPoints)
}

func (node *Node) model() *models.Node {
	return &models.Node{
		Id:    node.id,
		Name:  node.content.Name(),
		Data:  node.content.Data(),
		Size:  node.node.Size(),
		Color: node.node.Color(),
		Shape: node.node.Shape(),
	}
}

func (*Node) Draw() {
	panic("not implemented exception")
}

func (node *Node) Update(c chan struct{}, nodes *[]*physics.Node, graph *physics.Graph) {
	node.point.Update(c, nodes, graph)
}

func (node *Node) Move(time float64) {
	node.point.Move(time)
}

func (node *Node) SetNodeId(id int) {
	node.id = id
}

func (node *Node) SetNodeName(name string) {
	node.content.SetName(name)
}

func (node *Node) SetNodeData(data []byte) {
	node.content.SetData(data)
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
