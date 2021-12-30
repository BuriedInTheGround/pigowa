package main

import (
	"fmt"

	. "github.com/BuriedInTheGround/pigowa"
)

func main() {
	done := make(chan struct{})
	fmt.Println("Hello Gopher!")

	Setup(func() interface{} {
		CreateCanvas(400, 400)
		return nil
	})

	var ellipseSize float64 = 50

	Draw(func() interface{} {
		BackgroundGray(240)

		NoStroke()
		FillRGBA(212, 253, 121)
		Ellipse(200, 200, ellipseSize, ellipseSize)

		FillRGBA(212, 121, 253)
		Rect(0, 0, ellipseSize, ellipseSize, 20, 50, 200)

		FillRGBA(121, 212, 253)
		StrokeGray(120)
		StrokeWeight(5.34)
		Arc(200, 200, ellipseSize, ellipseSize, 0, PI+QUARTER_PI, PIE)

		ellipseSize = float64(int(ellipseSize+2) % 380)
		return nil
	})

	<-done
}
