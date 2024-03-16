package physics

import (
	"time"
)

type Graph struct {
	lastTime time.Time

	centralForceCoefficient   float64
	repulsiveForceCoefficient float64
	edgeLength                float64
	edgeStiffness             float64

	nodes *[]Node
}

func NewGraph(centralForceCoefficient, repulsiveForceCoefficient, edgeLength, edgeStiffness float64) *Graph {
	return &Graph{
		lastTime:                  time.Now(),
		centralForceCoefficient:   centralForceCoefficient,
		repulsiveForceCoefficient: repulsiveForceCoefficient,
		edgeLength:                edgeLength,
		edgeStiffness:             edgeStiffness,
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
