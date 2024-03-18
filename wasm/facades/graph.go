package facades

import (
	"graph-view-project/models"
	"graph-view-project/wasm/content"
	"graph-view-project/wasm/gui"
)

type Graph struct {
	id      int
	content content.Graph
	nodes   []*Node
	edges   []*Edge
}

func NewGraph(model models.Graph) *Graph {
	return &Graph{
		id: model.Id,
		content: *content.NewGraph(
			model.Name,
			model.Description,
		),
		nodes: make([]*Node, 0),
		edges: make([]*Edge, 0),
	}
}

func LoadGraph(model models.Graph) *Graph {
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
		nodes[node.Id] = newNode(node)
	}

	for _, edge := range model.Edges {
		edges[edge.Id] = newEdge(edge, nodes[edge.Begin], nodes[edge.End])
	}

	// TODO: add edges to Node.edges and Node.point.adjacentNodes

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

func DeleteNode(node *Node) {
	// TODO: delete node and all its edges
}

func DeleteEdge(edge *Edge) {
	// TODO: delete edge
}

func (graph *Graph) GetModel() *models.Graph {
	nodes := make([]*models.Node, len(graph.nodes))
	for i, node := range graph.nodes {
		nodes[i] = node.model()
	}

	edges := make([]*models.Edge, len(graph.edges))
	for i, edge := range graph.edges {
		edges[i] = edge.GetModel()
	}
	return models.NewGraph(
		graph.id,
		graph.content.Name(),
		graph.content.Description(),
		nodes,
		edges,
		gui.DefaultNode.Size(),
		gui.DefaultNode.Color(),
		gui.DefaultNode.Shape(),
		gui.DefaultEdge.Width(),
		gui.DefaultEdge.Color(),
		gui.DefaultEdge.Shape(),
	)
}
