// +build js, wasm

package viewtools

import (
	"github.com/josephbudd/crud/rendererprocess/framework/callback"
)

// Quit closes the application renderer.
func Quit() {
	callback.CloseCallBacks()
	global.Call("close")
}

// quit closes the application renderer.
func quit() {
	callback.CloseCallBacks()
	global.Call("close")
}
