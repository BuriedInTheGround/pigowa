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

	ellipseSize := 50

	Draw(func() interface{} {
		BackgroundGray(240)
		FillRGBA(121, 212, 253)
		Ellipse(200, 200, ellipseSize, ellipseSize)
		ellipseSize = (ellipseSize + 2) % 380
		return nil
	})

	<-done
}
