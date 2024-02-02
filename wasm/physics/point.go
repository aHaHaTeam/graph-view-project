package physics

type Point struct {
	isPinned     bool
	mass         float32
	position     Vec2d
	velocity     Vec2d
	acceleration Vec2d

	segments []Segment
}

func (p *Point) Update() {
	panic("not implemented exception")
}

func (p *Point) IsPinned() bool {
	return p.isPinned
}

func (p *Point) Pin() {
	p.isPinned = true
}

func (p *Point) UnPin() {
	p.isPinned = false
}
