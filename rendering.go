// Copyright 2021 Simone Ragusa
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pigowa

import "syscall/js"

func CreateCanvas(width, height int) {
	js.Global().Call("createCanvas", width, height)
	Width = width
	Height = height
}

func ResizeCanvas(width, height int) {
	js.Global().Call("resizeCanvas", width, height)
	Width = width
	Height = height
}
