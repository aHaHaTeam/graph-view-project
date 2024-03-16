package models

import "image/color"

type NodeShape int

const (
	Circle NodeShape = iota
	Square
)

type Node struct {
	Id    int
	Name  string
	Data  []byte
	Color color.Color
	Shape NodeShape
	Size  float32
}

// This seems to be VERY deprecated code. These functions were all moved to responsibilities of a database
/*
func (node *Node) Insert(db *sql.DB, color color.Color, shape gui.NodeShape, size float32) {
	data := string(node.data[:])
	_, err := db.Exec("INSERT INTO Nodes (id, name, data, color, shape, size) VALUES ($1, $2, $3, $4, $5, $6)",
		node.id, node.name, data, color, shape, size)
	if err != nil {
		log.Fatal(err)
	}
}

func (node *Node) Update(db *sql.DB, color color.Color, shape gui.NodeShape, size float32) {
	data := string(node.data[:])
	_, err := db.Exec("UPDATE Nodes SET name = $2, data = $3, color = $4, shape = $5, size = $6 WHERE id = $1",
		node.id, node.name, data, color, shape, size)
	if err != nil {
		log.Fatal(err)
	}
}

func (node *Node) Delete(db *sql.DB) {
	_, err := db.Exec("DELETE FROM Nodes WHERE id = $1", node.id)
	if err != nil {
		log.Fatal(err)
	}
}
*/
