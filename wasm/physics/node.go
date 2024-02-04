package physics

type Node struct {
	isPinned     bool
	mass         float32
	position     Vec2d
	velocity     Vec2d
	acceleration Vec2d

	edgesIn  []*Edge
	edgesOut []*Edge
}

func (p *Node) Update() {
	panic("not implemented exception")
}

func (p *Node) IsPinned() bool {
	return p.isPinned
}

func (p *Node) Pin() {
	p.isPinned = true
}

func (p *Node) UnPin() {
	p.isPinned = false
}
