package models

import "image/color"

type EdgeShape int

const (
	Arrow EdgeShape = iota
	Line
)

func EdgeShapeFromInt(value int) EdgeShape {
	switch value {
	case 0:
		return Arrow
	case 1:
		return Line
	default:
		return Arrow
	}
}

func EdgeShapeToInt(value EdgeShape) int {
	switch value {
	case Arrow:
		return 0
	case Line:
		return 1
	default:
		return 0
	}
}

type EdgeId int

type Edge struct {
	Id    EdgeId `json:"id"`
	Begin int    `json:"begin"`
	End   int    `json:"end"`

	Name        string `json:"name"`
	Description string `json:"description"`

	Width float32    `json:"width"`
	Color color.RGBA `json:"color"`
	Shape EdgeShape  `json:"shape"`
}

func NewEdge(
	id EdgeId,
	begin int,
	end int,
	name string,
	description string,
	width float32,
	color color.RGBA,
	shape EdgeShape,
) *Edge {
	return &Edge{
		Id:          id,
		Begin:       begin,
		End:         end,
		Name:        name,
		Description: description,
		Width:       width,
		Color:       color,
		Shape:       shape,
	}
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
