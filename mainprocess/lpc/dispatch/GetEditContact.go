package dispatch

import (
	"context"
	"fmt"

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

// handleGetEditContact is the *message.GetEditContactRendererToMainProcess handler.
// It's response back to the renderer is the *message.GetEditContactMainProcessToRenderer.
// Param ctx is the context. if <-ctx.Done() then the main process is shutting down.
// Param rxmessage *message.GetEditContactRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.GetEditContactMainProcessToRenderer message back to the renderer.
// Param stores is a struct the contains each of your stores.
// Param errChan is the channel to send the handler's error through since the handler does not return it's error.
func handleGetEditContact(ctx context.Context, rxmessage *message.GetEditContactRendererToMainProcess, sending lpc.Sending, stores *store.Stores, errChan chan error) {
	txmessage := &message.GetEditContactMainProcessToRenderer{
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
	if r == nil {
		// Send the err to package main.
		err = fmt.Errorf("unable to fine contact for %d", txmessage.ID)
		errChan <- err
		// Send the error to the renderer.
		// A bolt database error is fatal.
		txmessage.Error = true
		txmessage.ErrorMessage = err.Error()
		sending <- txmessage
		return
	}
	txmessage.Record = r
	sending <- txmessage
	return
}
