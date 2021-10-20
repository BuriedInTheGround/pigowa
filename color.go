// Copyright 2021 Simone Ragusa
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pigowa

import "syscall/js"

func BackgroundGray(color int, alpha ...int) {
	if len(alpha) == 0 {
		js.Global().Call("background", color)
	} else {
		js.Global().Call("background", color, alpha[0])
	}
}

func BackgroundRGBA(r, g, b int, alpha ...int) {
	if len(alpha) == 0 {
		js.Global().Call("background", r, g, b)
	} else {
		js.Global().Call("background", r, g, b, alpha[0])
	}
}

func BackgroundHex(code string) {
	js.Global().Call("background", code)
}

func Clear() {
	js.Global().Call("clear")
}

func FillGray(color int, alpha ...int) {
	if len(alpha) == 0 {
		js.Global().Call("fill", color)
	} else {
		js.Global().Call("fill", color, alpha[0])
	}
}

func FillRGBA(r, g, b int, alpha ...int) {
	if len(alpha) == 0 {
		js.Global().Call("fill", r, g, b)
	} else {
		js.Global().Call("fill", r, g, b, alpha[0])
	}
}

func FillHex(code string) {
	js.Global().Call("fill", code)
}

func NoFill() {
	js.Global().Call("noFill")
}

func StrokeGray(color int, alpha ...int) {
	if len(alpha) == 0 {
		js.Global().Call("stroke", color)
	} else {
		js.Global().Call("stroke", color, alpha[0])
	}
}

func StrokeRGBA(r, g, b int, alpha ...int) {
	if len(alpha) == 0 {
		js.Global().Call("stroke", r, g, b)
	} else {
		js.Global().Call("stroke", r, g, b, alpha[0])
	}
}

func StrokeHex(code string) {
	js.Global().Call("stroke", code)
}

func NoStroke() {
	js.Global().Call("noStroke")
}
