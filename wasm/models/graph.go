package models

import "database/sql"

type Graph struct {
	id          int
	name        string
	description string
	nodes       []Node
	edges       []Edge
}

func (graph *Graph) GetNodes(db *sql.DB) {
	//TODO
}
