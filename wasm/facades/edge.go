package facades

import (
	"graph-view-project/models"
	"graph-view-project/wasm/content"
	"graph-view-project/wasm/gui"
)

type Edge struct {
	id      models.EdgeId
	begin   *Node
	end     *Node
	content *content.Edge
	edge    *gui.Edge
}

func newEdge(model *models.Edge, begin, end *Node) *Edge {
	return &Edge{
		id:    model.Id,
		begin: begin,
		end:   end,
		content: content.NewEdge(
			model.Name,
			model.Description,
		),
		edge: gui.NewEdge(
			model.Width,
			model.Color,
			model.Shape,
		),
	}
}

func (edge *Edge) GetModel() *models.Edge {
	return &models.Edge{
		Id:          edge.id,
		Name:        edge.content.Name(),
		Description: edge.content.Description(),
		Width:       edge.edge.Width(),
		Color:       edge.edge.Color(),
		Shape:       edge.edge.Shape(),
	}
}

func (edge *Edge) SetBegin(node *Node) {
	edge.begin = node
}

func (edge *Edge) SetEnd(node *Node) {
	edge.end = node
}

func (this *Edge) Draw(canvas gui.Canvas) {
	x1, y1 := this.begin.Position()
	x2, y2 := this.end.Position()
	canvas.DrawEdge(this.edge, x1, y1, x2, y2)
}
