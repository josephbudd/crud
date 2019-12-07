package dispatch

import (
	"context"

	"github.com/josephbudd/crud/domain/lpc/message"
	"github.com/josephbudd/crud/domain/store"
	"github.com/josephbudd/crud/mainprocess/lpc"
)

/*
	YOU MAY EDIT THIS FILE.

	Rekickwasm will preserve this file for you.
	Kicklpc will not edit this file.

*/

// handleRemoveContact is the *message.RemoveContactRendererToMainProcess handler.
// It's response back to the renderer is the *message.RemoveContactMainProcessToRenderer.
// Param ctx is the context. if <-ctx.Done() then the main process is shutting down.
// Param rxmessage *message.RemoveContactRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.RemoveContactMainProcessToRenderer message back to the renderer.
// Param stores is a struct the contains each of your stores.
// Param errChan is the channel to send the handler's error through since the handler does not return it's error.
func handleRemoveContact(ctx context.Context, rxmessage *message.RemoveContactRendererToMainProcess, sending lpc.Sending, stores *store.Stores, errChan chan error) {
	txmessage := &message.RemoveContactMainProcessToRenderer{
		ID: rxmessage.ID,
	}
	if err := stores.Contact.Remove(rxmessage.ID); err != nil {
		txmessage.Error = true
		txmessage.ErrorMessage = err.Error()
	}
	sending <- txmessage
	// Also send the message that the contacts store has been modified.
	sending <- &message.ReloadContactsMainProcessToRenderer{}
	return
}
