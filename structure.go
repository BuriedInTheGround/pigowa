// Copyright 2021 Simone Ragusa
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pigowa

import "syscall/js"

// Setup executes the given function. It must be called only once, usually at
// the start of the program.
//
// It should be used to define initial environment properties such as screen
// size and background color.
func Setup(setup func() interface{}) {
	setupFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		initConstants()
		return setup()
	})
	js.Global().Set("setup", setupFunc)

	// TODO: maybe programmatically return an error if Setup is called more
	// than once(?)
}

// Draw continuously executes the given function until the program is stopped.
// It must be called only once, after Setup.
func Draw(draw func() interface{}) {
	drawFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return draw()
	})
	js.Global().Set("draw", drawFunc)
}
