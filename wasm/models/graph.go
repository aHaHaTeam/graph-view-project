package models

type Graph struct {
	id          int
	name        string
	description string
	nodes       []Node
	edges       []Edge
}
