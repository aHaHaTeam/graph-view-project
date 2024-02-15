package main

import (
	"fmt"
	. "github.com/BuriedInTheGround/pigowa"
	_ "github.com/joho/godotenv/autoload"
	"graph-view-project/wasm/database"
)

func main() {
	done := make(chan struct{})
	fmt.Println("Hello Gopher!")
	db, _ := database.Connect("graph-view-project")
	_ = db
	Setup(func() interface{} {
		//canvasSize := js.Global().Get("document").Call("getElementById", "canvas").Call("getBoundingClientRect")
		//CreateCanvas(canvasSize.Get("width").Int(), canvasSize.Get("height").Int())
		CreateCanvas(WindowWidth(), WindowHeight())
		return nil
	})

	doDraw := true

	ellipseSize := 50
	colors := []string{
		"#bf616a",
		"#8fbcbb",
		"#d08770",
		"#88c0d0",
		"#ebcb8b",
		"#81a1c1",
		"#a3be8c",
		"#5e81ac",
		"#b48ead",
	}
	colorIndex := 0

	Draw(func() interface{} {
		BackgroundRGBA(46, 52, 64, 100)
		StrokeWeight(3)
		StrokeHex("#d8dee9")
		FillHex(colors[colorIndex])
		if doDraw {
			if MouseIsPressed() && MouseButton() == "center" {
				StrokeWeight(10)
				Ellipse(float64(Width/2), float64(Height/2), float64(ellipseSize), float64(ellipseSize))
			} else {
				Ellipse(float64(MouseX()), float64(MouseY()), float64(ellipseSize), float64(ellipseSize))
			}
		}
		return nil
	})

	WindowResized(func() interface{} {
		ResizeCanvas(WindowWidth(), WindowHeight())
		return nil
	})

	MousePressed(func() interface{} {
		if MouseButton() == "right" {
			doDraw = !doDraw
		}
		if !doDraw {
			Clear()
		}
		return false
	})

	MouseClicked(func() interface{} {
		colorIndex = (colorIndex + 1) % len(colors)
		return false
	})

	MouseWheel(func(delta float64) interface{} {
		ellipseSize -= int(delta * 0.05)
		return false
	})

	<-done
}
