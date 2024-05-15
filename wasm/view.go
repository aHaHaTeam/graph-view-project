package main

import (
	"syscall/js"

	. "github.com/aHaHaTeam/p5js-wasm-go/color"
	. "github.com/aHaHaTeam/p5js-wasm-go/events"
	. "github.com/aHaHaTeam/p5js-wasm-go/p5js"
	. "github.com/aHaHaTeam/p5js-wasm-go/rendering"
	. "github.com/aHaHaTeam/p5js-wasm-go/shape"
)

func startDrawing() {
	radius := 20.0
	colorIndex := 0
	colors := [][]any{{239, 30, 30}, {30, 239, 30}, {30, 30, 239}}

	Setup(func(...js.Value) interface{} {
		_, err := CreateCanvas(400, 400)
		if err != nil {
			return nil
		}
		err = Background(100)
		if err != nil {
			return nil
		}
		return nil
	})

	Draw(func(...js.Value) interface{} {
		err := Ellipse(MouseX(), MouseY(), radius, radius)
		if err != nil {
			return nil
		}
		return nil
	})

	MouseMoved(func(args ...js.Value) interface{} {
		_ = NoStroke()
		_ = Fill(colors[colorIndex%len(colors)]...)
		return nil
	})

	MouseClicked(func(args ...js.Value) interface{} {
		colorIndex = colorIndex + 1
		return nil
	})

	MouseDragged(func(args ...js.Value) interface{} {
		_ = Stroke(255, 255, 255)
		_ = Fill(255, 255, 255, 63)
		_ = Ellipse(MouseX(), MouseY(), radius, radius)
		return nil
	})

	MouseWheel(func(args ...js.Value) interface{} {
		if args[0].Get("delta").Int() > 0 {
			radius *= 1.1
		} else {
			radius /= 1.1
		}
		return nil
	})
	select {}
}
