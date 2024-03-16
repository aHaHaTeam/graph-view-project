package models

import "image/color"

type EdgeShape int

const (
	Arrow EdgeShape = iota
	Line
)

type Edge struct {
	Id          int
	Name        string
	Description string
	Width       float32
	Color       color.Color
	Shape       EdgeShape
}

// This seems to be VERY deprecated code. These functions were all moved to responsibilities of a database
/*
func (edge *Edge) Insert(db *sql.DB, color color.Color, shape gui.EdgeShape, width float32) {
	data := string(edge.description[:])
	_, err := db.Exec("INSERT INTO Edges (id, name, data, color, shape, width) VALUES ($1, $2, $3, $4, $5, $6)",
		edge.id, edge.name, data, color, shape, width)
	if err != nil {
		log.Fatal(err)
	}
}

func (edge *Edge) Update(db *sql.DB, color color.Color, shape gui.NodeShape, size float32) {
	data := string(edge.description[:])
	_, err := db.Exec("UPDATE Edges SET name = $2, data = $3, color = $4, shape = $5, size = $6 WHERE id = $1",
		edge.id, edge.name, data, color, shape, size)
	if err != nil {
		log.Fatal(err)
	}
}

func (edge *Edge) Delete(db *sql.DB) {
	_, err := db.Exec("DELETE FROM Edges WHERE id = $1", edge.id)
	if err != nil {
		log.Fatal(err)
	}
}
*/
