package facades

import (
	"database/sql"
	"graph-view-project/models"
	"log"
)

type Graph struct {
	model models.Graph
	Nodes []Node
	Edges []Edge
}

func (graph *Graph) GetId() int {
	return graph.model.Id
}

func (graph *Graph) LoadNodes(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM Nodes WHERE id = $1", graph.GetId())
	if err != nil {
		log.Fatal(err)
	}

	graph.Nodes = make([]Node, 0)
	for rows.Next() {
		node := Node{}
		err = rows.Scan(node.SetNodeId, node.SetNodeName, node.SetNodeData,
			node.SetNodeColor, node.SetNodeShape, node.SetNodeSize)
		graph.Nodes = append(graph.Nodes, node)
	}
}
