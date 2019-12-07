package dispatch

import (
	"context"

	"github.com/josephbudd/crud/domain/lpc/message"
	"github.com/josephbudd/crud/domain/store"
	"github.com/josephbudd/crud/domain/store/record"
	"github.com/josephbudd/crud/mainprocess/lpc"
)

/*
	YOU MAY EDIT THIS FILE.

	Rekickwasm will preserve this file for you.
	Kicklpc will not edit this file.

*/

// handleGetPrintContact is the *message.GetPrintContactRendererToMainProcess handler.
// It's response back to the renderer is the *message.GetPrintContactMainProcessToRenderer.
// Param ctx is the context. if <-ctx.Done() then the main process is shutting down.
// Param rxmessage *message.GetPrintContactRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.GetPrintContactMainProcessToRenderer message back to the renderer.
// Param stores is a struct the contains each of your stores.
// Param errChan is the channel to send the handler's error through since the handler does not return it's error.
func handleGetPrintContact(ctx context.Context, rxmessage *message.GetPrintContactRendererToMainProcess, sending lpc.Sending, stores *store.Stores, errChan chan error) {
	txmessage := &message.GetPrintContactMainProcessToRenderer{
		ID: rxmessage.ID,
	}
	var r *record.Contact
	var err error
	if r, err = stores.Contact.Get(rxmessage.ID); err != nil {
		txmessage.Error = true
		txmessage.ErrorMessage = err.Error()
		sending <- txmessage
		return
	}
	txmessage.Record = r
	sending <- txmessage
	return
}
