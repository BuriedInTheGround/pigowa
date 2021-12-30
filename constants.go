package pigowa

import (
	"syscall/js"
)

var (
	HALF_PI    float64
	PI         float64
	QUARTER_PI float64
	TAU        float64
	TWO_PI     float64
	DEGREES    js.Value
	RADIANS    js.Value

	CHORD js.Value
	OPEN  js.Value
	PIE   js.Value
)

func initConstants() {
	HALF_PI = js.Global().Get("HALF_PI").Float()
	PI = js.Global().Get("PI").Float()
	QUARTER_PI = js.Global().Get("QUARTER_PI").Float()
	TAU = js.Global().Get("TAU").Float()
	TWO_PI = js.Global().Get("TWO_PI").Float()
	DEGREES = js.Global().Get("DEGREES")
	RADIANS = js.Global().Get("RADIANS")

	CHORD = js.Global().Get("CHORD")
	OPEN = js.Global().Get("OPEN")
	PIE = js.Global().Get("PIE")
}
