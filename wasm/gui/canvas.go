package gui

import (
	"github.com/aHaHaTeam/p5js-wasm-go/color"
	"github.com/aHaHaTeam/p5js-wasm-go/rendering"
	"github.com/aHaHaTeam/p5js-wasm-go/shape"
	"log"
	"syscall/js"
)

type Canvas struct {
	width, height float32
	jsCanvas      js.Value
	dx, dy        float32
	zoom          float32
}

func NewCanvas(width, height float32) *Canvas {
	canvas, err := rendering.CreateCanvas(width, height)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &Canvas{width: width, height: height, jsCanvas: canvas}
}

func (canvas *Canvas) toScreenSpace(x, y float32) (float32, float32) {
	return (x-canvas.dx)*canvas.zoom + canvas.width/2,
		(y-canvas.dy)*canvas.zoom + canvas.height/2
}

func (canvas *Canvas) DrawNode(node *Node, x, y float32) {
	_ = color.Stroke(node.color.RGBA())
	screenX, screenY := canvas.toScreenSpace(x, y)
	_ = shape.Circle(screenX, screenY, node.Size())
}

func (canvas *Canvas) DrawEdge(edge *Edge, x1, y1 float32, x2, y2 float32) {
	_ = color.Stroke(edge.color.RGBA())
	// TODO add strokeWeight
	screenX1, screenY1 := canvas.toScreenSpace(x1, y1)
	screenX2, screenY2 := canvas.toScreenSpace(x2, y2)
	_ = shape.Line(screenX1, screenY1, screenX2, screenY2)
}
