// Copyright 2021 Simone Ragusa
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pigowa

import (
	"fmt"
	"syscall/js"
)

// Arc draws an arc on the canvas.
//
// If called with only x, y, width, height, start and stop, the arc will be
// drawn and filled as an open pie segment. If a mode parameter is provided,
// the arc will be filled like an open semi-circle (OPEN), a closed semi-circle
// (CHORD), or as a closed pie segment (PIE).
func Arc(x, y, width, height, start, stop float64, mode ...js.Value) error {
	if len(mode) > 1 {
		return fmt.Errorf(
			"is accepted at most one value for mode, %d given", len(mode),
		)
	}

	if len(mode) == 0 {
		js.Global().Call("arc", x, y, width, height, start, stop)
	} else {
		js.Global().Call("arc", x, y, width, height, start, stop, mode[0].String())
	}

	return nil
}

// Ellipse draws an ellipse (oval) to the canvas.
//
// The first two parameters set the location of the center of the ellipse, and
// the third and fourth parameters set the shape's width and height.
func Ellipse(x, y, width, height float64) {
	js.Global().Call("ellipse", x, y, width, height)
}

// Rect draws a rectangle on the canvas.
//
// The first two parameters set the location of the upper-left corner, the
// third sets the width, and the fourth sets the height. The optional parameter
// specifies the corner radius for the top-left, top-right, lower-right and
// lower-left corners, respectively. An omitted corner radius parameter is set
// to the value of the previously specified radius value in the parameter list.
func Rect(x, y, width, height float64, cornerRadius ...float64) error {
	// Get how many corner radius parameters we receive.
	crlen := len(cornerRadius)

	// If too many, return early with error.
	if crlen > 4 {
		return fmt.Errorf(
			"are accepted at most four values for cornerRadius, %d given",
			len(cornerRadius),
		)
	}

	// If zero, set all to be zero.
	if crlen == 0 {
		cornerRadius = []float64{0, 0, 0, 0}
		crlen = 4
	}

	// If less than four, fill the others with the value of the last one given.
	for i := crlen; i < 4; i++ {
		cornerRadius = append(cornerRadius, cornerRadius[crlen-1])
	}

	js.Global().Call(
		"rect", x, y, width, height,
		cornerRadius[0], cornerRadius[1], cornerRadius[2], cornerRadius[3],
	)

	return nil
}

// StrokeWeight sets the width of the stroke used for lines, points and the
// border around shapes.
func StrokeWeight(weight float64) {
	js.Global().Call("strokeWeight", weight)
}
