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

// handleGetRemoveContact is the *message.GetRemoveContactRendererToMainProcess handler.
// It's response back to the renderer is the *message.GetRemoveContactMainProcessToRenderer.
// Param ctx is the context. if <-ctx.Done() then the main process is shutting down.
// Param rxmessage *message.GetRemoveContactRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.GetRemoveContactMainProcessToRenderer message back to the renderer.
// Param stores is a struct the contains each of your stores.
// Param errChan is the channel to send the handler's error through since the handler does not return it's error.
func handleGetRemoveContact(ctx context.Context, rxmessage *message.GetRemoveContactRendererToMainProcess, sending lpc.Sending, stores *store.Stores, errChan chan error) {
	txmessage := &message.GetRemoveContactMainProcessToRenderer{
		ID: rxmessage.ID,
	}
	var r *record.Contact
	var err error
	if r, err = stores.Contact.Get(rxmessage.ID); err != nil {
		// Send the err to package main.
		errChan <- err
		// Send the error to the renderer.
		// A bolt database error is fatal.
		txmessage.Fatal = true
		txmessage.ErrorMessage = err.Error()
		sending <- txmessage
		return
	}
	txmessage.Record = r
	sending <- txmessage
	return
}
