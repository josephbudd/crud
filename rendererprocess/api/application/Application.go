// +build js, wasm

package application

import (
	"context"

	"github.com/josephbudd/crud/rendererprocess/api/display"
	"github.com/josephbudd/crud/rendererprocess/api/event"
	"github.com/josephbudd/crud/rendererprocess/framework/callback"
)

// NewGracefullyCloseHandler makes an event handler which gracefully closes the application for you.
// Use it in your panel controllers to handle your own application closing buttons.
// Param cancelFunc is the renderer process's context cancel func.
func NewGracefullyCloseHandler(cancelFunc context.CancelFunc) (handler func(e event.Event) (nilReturn interface{})) {
	handler = func(e event.Event) (nilReturn interface{}) {
		GracefullyClose(cancelFunc)
		return
	}
	return
}

// GracefullyClose gracefully closes the application for you.
// Use it in your panel controllers.
// Param cancelFunc is the renderer process's context cancel func.
func GracefullyClose(cancelFunc context.CancelFunc) {
	callback.RemoveApplicationOnCloseHandler()
	title := "Closing"
	msg := "Closing <q>Spawn Tabs</q>."
	display.Inform(msg, title, cancelFunc)
	return
}
