package models

type Graph struct {
	Id          int
	Name        string
	Description string
	Nodes       []Node
	Edges       []Edge
}
