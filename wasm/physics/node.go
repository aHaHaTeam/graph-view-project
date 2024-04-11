package physics

type Node struct {
	isPinned bool
	position Vec2d
	velocity Vec2d

	adjacentNodes []*Node
}

func NewNode() *Node {
	return &Node{
		isPinned:      false,
		position:      Vec2d{0, 0},
		velocity:      Vec2d{0, 0},
		adjacentNodes: make([]*Node, 0),
	}
}

func (node *Node) AddAdjacentNode(adjacentNode *Node) {
	node.adjacentNodes = append(node.adjacentNodes, adjacentNode)
}

func (node *Node) RemoveAdjacentNode(adjacentNode *Node) {
	for i, n := range node.adjacentNodes {
		if n == adjacentNode {
			node.adjacentNodes = append(node.adjacentNodes[:i], node.adjacentNodes[i+1:]...)
			break
		}
	}
}

func (node *Node) Update(c chan struct{}, nodes *[]*Node, graph *Graph) {
	if node.isPinned {
		node.velocity = Vec2d{0, 0}
		return
	}

	resultantForce := centralForce(node, graph)

	for _, n := range node.adjacentNodes {
		resultantForce.add(springForce(node, n, graph))
	}

	for _, n := range *nodes {
		resultantForce.add(repulsiveForce(node, n, graph))
	}

	node.velocity.add(resultantForce.divide(graph.parameters.NodeMass))

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
