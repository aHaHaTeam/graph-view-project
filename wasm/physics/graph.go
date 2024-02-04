package physics

import (
	"graph-view-project/wasm/facades"
	"time"
)

type Graph struct {
	lastTime time.Time
	//Some useful constants
}

func (g *Graph) Update(nodes *[]facades.Node) {
	c := make(chan struct{})
	for _, n := range *nodes {
		go n.Update(c)
	}

	for range *nodes {
		<-c
	}

	dt := time.Since(g.lastTime).Seconds()
	g.lastTime = time.Now()
	for _, n := range *nodes {
		n.Move(dt)
	}
}
