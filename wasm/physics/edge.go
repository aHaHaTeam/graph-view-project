package physics

type Edge struct {
	begin, end *Node
	length     float64
	k          float64
}

func (edge *Edge) GetLength() float64 {
	vec := vectorFromTwoPoints(edge.begin.position, edge.end.position)
	return vec.length()
}
