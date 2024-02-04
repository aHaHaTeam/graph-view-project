package physics

type Node struct {
	isPinned bool
	mass     float64
	position Vec2d
	velocity Vec2d

	edgesIn  []*Edge
	edgesOut []*Edge
}

func (p *Node) Update(c chan struct{}) {
	if p.isPinned {
		return
	}

	resultantForce := Vec2d{0, 0}

	for _, e := range p.edgesIn {
		v := e.begin.position.subtract(e.end.position)
		dl := v.length() - e.length
		resultantForce.add(v.multiply(e.k * dl))
	}

	for _, e := range p.edgesIn {
		v := e.end.position.subtract(e.begin.position)
		dl := v.length() - e.length
		resultantForce.add(v.multiply(e.k * dl))
	}

	p.velocity.add(resultantForce.divide(p.mass))

	c <- struct{}{}
}

func (p *Node) Move(time float64) {
	p.position.add(p.velocity.multiply(time))
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
