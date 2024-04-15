package physics

import (
	"time"
)

type PhysicalParameters struct {
	CentralForceCoefficient   float64
	RepulsiveForceCoefficient float64
	EdgeLength                float64
	EdgeStiffness             float64
	NodeMass                  float64
}

type Graph struct {
	parameters PhysicalParameters
	nodes      *[]*Node

	lastTime time.Time
}

func NewGraph(parameters PhysicalParameters, nodes *[]*Node) *Graph {
	return &Graph{
		parameters: parameters,
		nodes:      nodes,
		lastTime:   time.Now(),
	}
}

func (graph *Graph) Update() {
	c := make(chan struct{})
	for _, n := range *graph.nodes {
		go n.Update(c, graph.nodes, graph)
	}

	for range *graph.nodes {
		<-c
	}

	dt := time.Since(graph.lastTime).Seconds()
	graph.lastTime = time.Now()
	for _, n := range *graph.nodes {
		n.Move(dt)
	}
}
