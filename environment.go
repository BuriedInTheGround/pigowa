// Copyright 2021 Simone Ragusa
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pigowa

import "syscall/js"

var (
	// Width of the drawing canvas. Defaults to 100. Updates upon use.
	Width = 100

	// Height of the drawing canvas. Defaults to 100. Updates upon use.
	Height = 100
)

// WindowWidth returns the width of the inner window, it maps to
// `window.innerWidth`.
func WindowWidth() int {
	return js.Global().Get("windowWidth").Int()
}

// WindowHeight returns the height of the inner window, it maps to
// `window.innerHeight`.
func WindowHeight() int {
	return js.Global().Get("windowHeight").Int()
}

// WindowResized executes the given function once every time the browser window
// is resized.
//
// This is a good place to resize the canvas or do any other adjustments to
// accommodate the new window size.
func WindowResized(resized func() interface{}) {
	resizedFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return resized()
	})
	js.Global().Set("windowResized", resizedFunc)
}
