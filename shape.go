// Copyright 2021 Simone Ragusa
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pigowa

import "syscall/js"

func Ellipse(x, y, width, height int) {
	js.Global().Call("ellipse", x, y, width, height)
}

func StrokeWeight(weight int) {
	js.Global().Call("strokeWeight", weight)
}
