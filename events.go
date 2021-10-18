// Copyright 2021 Simone Ragusa
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pigowa

import "syscall/js"

func MouseX() int {
	return js.Global().Get("mouseX").Int()
}

func MouseY() int {
	return js.Global().Get("mouseY").Int()
}

func MouseClicked(clicked func() interface{}) {
	clickedFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return clicked()
	})
	js.Global().Set("mouseClicked", clickedFunc)
}

func MouseWheel(wheel func(delta float64) interface{}) {
	wheelFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			return wheel(args[0].Get("delta").Float())
		} else {
			panic("pigowa: missing event argument expected")
		}
	})
	js.Global().Set("mouseWheel", wheelFunc)
}
