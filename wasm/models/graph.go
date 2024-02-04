package models

import (
	"database/sql"
	"graph-view-project/wasm/facades"
	"log"
)

type Graph struct {
	id          int
	name        string
	description string
}

func (graph *Graph) GetNodes(db *sql.DB, nodes *[]facades.Node) {
	rows, err := db.Query("SELECT * FROM Nodes WHERE id = $1", graph.id)
	if err != nil {
		log.Fatal(err)
	}

	*nodes = make([]facades.Node, 0)
	for rows.Next() {
		node := facades.Node{}
		err = rows.Scan(node.SetNodeId, node.SetNodeName, node.SetNodeData,
			node.SetNodeColor, node.SetNodeShape, node.SetNodeSize)
		*nodes = append(*nodes, node)
	}
}
