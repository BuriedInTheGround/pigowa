// Copyright 2021 Simone Ragusa
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pigowa

import "syscall/js"

var (
	Width  = 100
	Height = 100
)

func WindowWidth() int {
	return js.Global().Get("windowWidth").Int()
}

func WindowHeight() int {
	return js.Global().Get("windowHeight").Int()
}

func WindowResized(resized func() interface{}) {
	resizedFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return resized()
	})
	js.Global().Set("windowResized", resizedFunc)
}
