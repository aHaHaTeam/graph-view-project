package physics

func centralForce(node *Node, graph *Graph) Vec2d {
	r := node.position.length()
	return node.position.multiply(-graph.centralForceCoefficient * node.mass / (r * r * r))
}

func repulsiveForce(node1 *Node, node2 Node, graph *Graph) Vec2d {
	r := node1.position.subtract(node2.position).length()
	return node1.position.subtract(node2.position).multiply(
		graph.repulsiveForceCoefficient * node1.mass * node2.mass / (r * r * r))
}

func springForce(node1 *Node, node2 *Node, edge *Edge, graph *Graph) Vec2d {
	v := node2.position.subtract(node1.position)
	dl := v.length() - graph.edgeLength
	return v.multiply(graph.edgeStiffness * dl)
}
