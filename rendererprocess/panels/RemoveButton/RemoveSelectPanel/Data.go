// +build js, wasm

package removeselectpanel

import (
	"github.com/josephbudd/crud/rendererprocess/api/dom"
	"github.com/josephbudd/crud/rendererprocess/framework/lpc"
)

/*

	Panel name: RemoveSelectPanel

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

	// The document object module.
	document *dom.DOM

	sortedIndexAttributeName = "sortedIndex"
	recordIDAttributeName    = "recordID"
)
