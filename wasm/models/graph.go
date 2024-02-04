package models

import (
	"database/sql"
	"graph-view-project/wasm/facades"
)

type Graph struct {
	id          int
	name        string
	description string
}

func (graph *Graph) GetNodes(db *sql.DB, nodes *[]facades.Node, edges *[]facades.Edge) {
	//TODO
}
