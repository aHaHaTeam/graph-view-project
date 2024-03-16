package physics

type Edge struct {
	begin, end *Node
}

func NewEdge(begin, end *Node) *Edge {
	return &Edge{
		begin: begin,
		end:   end,
	}
}
