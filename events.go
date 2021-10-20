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

func MouseButton() string {
	return js.Global().Get("mouseButton").String()
}

func MouseIsPressed() bool {
	return js.Global().Get("mouseIsPressed").Bool()
}

func MousePressed(pressed func() interface{}) {
	pressedFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return pressed()
	})
	js.Global().Set("mousePressed", pressedFunc)
}

func MouseReleased(released func() interface{}) {
	releasedFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return released()
	})
	js.Global().Set("mouseReleased", releasedFunc)
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
