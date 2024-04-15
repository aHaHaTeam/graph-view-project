package physics

func centralForce(node *Node, graph *Graph) Vec2d {
	r := node.position.length()
	return node.position.multiply(
		-graph.parameters.CentralForceCoefficient *
			graph.parameters.NodeMass / (r * r * r))
}

func repulsiveForce(node1 *Node, node2 *Node, graph *Graph) Vec2d {
	r := node1.position.subtract(node2.position).length()
	return node1.position.subtract(node2.position).multiply(
		graph.parameters.RepulsiveForceCoefficient *
			graph.parameters.NodeMass * graph.parameters.NodeMass / (r * r * r))
}

func springForce(node1 *Node, node2 *Node, graph *Graph) Vec2d {
	v := node2.position.subtract(node1.position)
	dl := v.length() - graph.parameters.EdgeLength
	return v.multiply(graph.parameters.EdgeStiffness * dl)
}
