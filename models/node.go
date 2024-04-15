package models

import "image/color"

type NodeShape int

const (
	Circle NodeShape = iota
	Square
)

func NodeShapeFromInt(value int) NodeShape {
	switch value {
	case 0:
		return Circle
	case 1:
		return Square
	default:
		return Circle
	}
}

func NodeShapeToInt(value NodeShape) int {
	switch value {
	case Circle:
		return 0
	case Square:
		return 1
	default:
		return 0
	}
}

type Node struct {
	Id    int   `json:"id"`
	Edges []int `json:"edges"`

	Name string `json:"name"`
	Data []byte `json:"data"`

	Size  float32    `json:"size"`
	Color color.RGBA `json:"color"`
	Shape NodeShape  `json:"shape"`
}

func NewNode(
	id int,
	edges []int,
	name string,
	data []byte,
	size float32,
	color color.RGBA,
	shape NodeShape,
) *Node {
	return &Node{
		Id:    id,
		Edges: edges,
		Name:  name,
		Data:  data,
		Size:  size,
		Color: color,
		Shape: shape,
	}
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
