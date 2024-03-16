package facades

import (
	"graph-view-project/models"
	"graph-view-project/wasm/content"
	"graph-view-project/wasm/gui"
	"graph-view-project/wasm/physics"
)

type Edge struct {
	id      int
	begin   *Node
	end     *Node
	content *content.Edge
	segment *physics.Edge
	edge    *gui.Edge
}

func NewEdge(model *models.Edge) *Edge {
	return &Edge{
		id: model.Id,

		content: content.NewEdge(
			model.Name,
			model.Description,
		),
		segment: &physics.Edge{},
		edge: gui.NewEdge(
			model.Width,
			model.Color,
			model.Shape,
		),
	}
}

func (e *Edge) GetModel() models.Edge {
	return models.Edge{
		Id:          e.id,
		Name:        e.content.Name(),
		Description: e.content.Description(),
		Width:       e.edge.Width(),
		Color:       e.edge.Color(),
		Shape:       e.edge.Shape(),
	}
}

func (e *Edge) SetBegin(node *Node) {
	e.begin = node
}

func (e *Edge) SetEnd(node *Node) {
	e.end = node
}

func (*Edge) Draw() {
	panic("not implemented exception")
}
