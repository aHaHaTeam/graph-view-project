package physics

import (
	"time"
)

type Graph struct {
	lastTime time.Time

	centralForceCoefficient   float64
	repulsiveForceCoefficient float64
}

func (graph *Graph) Update(nodes *[]Node) {
	c := make(chan struct{})
	for _, n := range *nodes {
		go n.Update(c, nodes, graph)
	}

	for range *nodes {
		<-c
	}

	dt := time.Since(graph.lastTime).Seconds()
	graph.lastTime = time.Now()
	for _, n := range *nodes {
		n.Move(dt)
	}
}
