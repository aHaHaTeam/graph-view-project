package physics

type Node struct {
	isPinned bool
	mass     float64
	position Vec2d
	velocity Vec2d

	edgesIn  []*Edge
	edgesOut []*Edge
}

func (node *Node) Update(c chan struct{}, nodes *[]Node, graph *Graph) {
	if node.isPinned {
		node.velocity = Vec2d{0, 0}
		return
	}

	resultantForce := centralForce(node, graph)

	for _, e := range node.edgesIn {
		resultantForce.add(springForce(e.end, e.begin, e))
	}

	for _, e := range node.edgesOut {
		resultantForce.add(springForce(e.begin, e.end, e))
	}

	for _, n := range *nodes {
		resultantForce.add(repulsiveForce(node, n, graph))
	}

	node.velocity.add(resultantForce.divide(node.mass))

	c <- struct{}{}
}

func (node *Node) Move(time float64) {
	node.position.add(node.velocity.multiply(time))
}

func (node *Node) IsPinned() bool {
	return node.isPinned
}

func (node *Node) Pin() {
	node.isPinned = true
}

func (node *Node) UnPin() {
	node.isPinned = false
}
