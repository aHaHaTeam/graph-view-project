package physics

type Edge struct {
	begin, end *Node
	length     float64
	k          float64
}

func (e *Edge) GetLength() float64 {
	vec := vectorFromTwoPoints(e.begin.position, e.end.position)
	return vec.length()
}
