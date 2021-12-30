// Copyright 2021 Simone Ragusa
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pigowa

import "syscall/js"

func Setup(setup func() interface{}) {
	setupFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		initConstants()
		return setup()
	})
	js.Global().Set("setup", setupFunc)
}

func Draw(draw func() interface{}) {
	drawFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return draw()
	})
	js.Global().Set("draw", drawFunc)
}
