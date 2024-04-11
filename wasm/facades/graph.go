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
		begin := nodes[edge.Begin]
		end := nodes[edge.End]
		facade := newEdge(edge, begin, end)

		edges[edge.Id] = facade

		begin.edges = append(begin.edges, facade)
		begin.point.AddAdjacentNode(end.point)

		end.edges = append(end.edges, facade)
		end.point.AddAdjacentNode(begin.point)
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

func (graph *Graph) DeleteNode(node *Node) {
	// delete actual node
	for i, n := range graph.nodes {
		if n.id == node.id {
			graph.nodes = append(graph.nodes[:i], graph.nodes[i+1:]...)
		}
	}

	// delete edges from deleted node

	remainingEdges := make([]*Edge, 0)
	// chose what edges to keep
	for _, e := range graph.edges {
		if e.begin.id == node.id {
			// remove the edge from another side
			e.end.removeEdge(e)
		} else if e.end.id == node.id {
			e.begin.removeEdge(e)
		} else {
			remainingEdges = append(remainingEdges, e)
		}

	}
	graph.edges = remainingEdges
}

func (graph *Graph) DeleteEdge(edge *Edge) {
	for i, e := range graph.edges {
		if e.id == edge.id {
			e.begin.removeEdge(e)
			e.end.removeEdge(e)
			graph.edges = append(graph.edges[:i], graph.edges[i+1:]...)
			break
		}
	}
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
