// Copyright 2021 Simone Ragusa
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pigowa

import "syscall/js"

// CreateCanvas creates a canvas element in the document, and sets the
// dimensions of it. The origin (0,0) is positioned at the top left of the
// screen.
//
// This method should be called only once at the start of setup. Calling
// CreateCanvas more than once will result in very unpredictable behavior.
func CreateCanvas(width, height int) {
	js.Global().Call("createCanvas", width, height)
	Width = width
	Height = height
}

// ResizeCanvas resizes the canvas to given width and height. The canvas will
// be cleared and draw will be called immediately, allowing the sketch to
// re-render itself in the resized canvas.
func ResizeCanvas(width, height int) {
	js.Global().Call("resizeCanvas", width, height)
	Width = width
	Height = height
}
