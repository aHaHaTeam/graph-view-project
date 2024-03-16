package facades

import (
	"graph-view-project/models"
	"graph-view-project/wasm/content"
)

type Graph struct {
	id      int
	content content.Graph
	nodes   []*Node
	edges   []*Edge
}

func NewGraph(model models.Graph) *Graph {
	nodes := make(map[int]*Node)
	edges := make(map[int]*Edge)

	graph := Graph{
		id: model.Id,
		content: *content.NewGraph(
			model.Name,
			model.Description,
		),
	}

	for _, node := range model.Nodes {
		nodes[node.Id] = NewNode(&node)
	}

	for _, edge := range model.Edges {
		edges[edge.Id] = NewEdge(&edge)
	}

	for _, node := range model.Nodes {
		for _, edgeId := range node.Edges {
			nodes[node.Id].edges = append(nodes[node.Id].edges, edges[edgeId])
		}
	}

	for _, edge := range model.Edges {
		edges[edge.Id].SetBegin(nodes[edge.Begin])
		edges[edge.Id].SetEnd(nodes[edge.End])
	}

	graph.nodes = make([]*Node, len(nodes))
	graph.edges = make([]*Edge, len(edges))

	for i, node := range nodes {
		graph.nodes[i] = node
	}

	for i, edge := range edges {
		graph.edges[i] = edge
	}

	return &graph
}

func (g *Graph) GetModel() models.Graph {
	nodes := make([]models.Node, len(g.nodes))
	for i, node := range g.nodes {
		nodes[i] = node.GetModel()
	}

	edges := make([]models.Edge, len(g.edges))
	for i, edge := range g.edges {
		edges[i] = edge.GetModel()
	}
	return models.Graph{
		Id:          g.id,
		Name:        g.content.Name(),
		Description: g.content.Description(),
		Nodes:       nodes,
		Edges:       edges,
	}
}
