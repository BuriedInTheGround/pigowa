// Copyright 2021 Simone Ragusa
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pigowa

import (
	"syscall/js"
)

func init() {
	document := js.Global().Get("document")

	p5 := document.Call("createElement", "script")
	p5.Set("src", "https://cdnjs.cloudflare.com/ajax/libs/p5.js/1.4.0/p5.js")

	styles := document.Call("createElement", "style")
	styles.Set("innerHTML", `
html, body {
    margin: 0;
    padding: 0;
}
canvas {
    display: block;
}`)

	document.Get("head").Call("appendChild", p5)
	document.Get("head").Call("appendChild", styles)
}
