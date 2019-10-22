// +build js, wasm

package editformpanel

import (
	"syscall/js"

	"github.com/josephbudd/crud/rendererprocess/lpc"
	"github.com/josephbudd/crud/rendererprocess/notjs"
	"github.com/josephbudd/crud/rendererprocess/viewtools"
)

/*

	Panel name: EditFormPanel

*/

var (
	// quitCh will close the application
	quitCh chan struct{}

	// eojCh will signal go routines to stop and return because the application is ending.
	eojCh chan struct{}

	// receiveCh receives messages from the main process.
	receiveCh lpc.Receiving

	// sendCh sends messages to the main process.
	sendCh lpc.Sending

	// The framework's renderer API.
	tools *viewtools.Tools

	// Some javascipt like dom functions written in go.
	notJS *notjs.NotJS

	// The javascript null value
	null = js.Null()
)
