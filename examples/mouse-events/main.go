package main

import (
	"fmt"

	. "github.com/BuriedInTheGround/pigowa"
)

func main() {
	done := make(chan struct{}, 0)
	fmt.Println("Hello Gopher!")

	Setup(func() interface{} {
		CreateCanvas(WindowWidth(), WindowHeight())
		return nil
	})

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
		Ellipse(MouseX(), MouseY(), ellipseSize, ellipseSize)
		return nil
	})

	WindowResized(func() interface{} {
		ResizeCanvas(WindowWidth(), WindowHeight())
		return nil
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