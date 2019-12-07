package dispatch

import (
	"context"
	"fmt"
	"log"

	"github.com/josephbudd/crud/domain/lpc/message"
	"github.com/josephbudd/crud/domain/store"
	"github.com/josephbudd/crud/mainprocess/lpc"
)

/*
	YOU MAY EDIT THIS FILE.

	Rekickwasm will preserve this file for you.
	Kicklpc will not edit this file.

*/

// handleEditContact is the *message.EditContactRendererToMainProcess handler.
// It's response back to the renderer is the *message.EditContactMainProcessToRenderer.
// Param ctx is the context. if <-ctx.Done() then the main process is shutting down.
// Param rxmessage *message.EditContactRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.EditContactMainProcessToRenderer message back to the renderer.
// Param stores is a struct the contains each of your stores.
// Param errChan is the channel to send the handler's error through since the handler does not return it's error.
func handleEditContact(ctx context.Context, rxmessage *message.EditContactRendererToMainProcess, sending lpc.Sending, stores *store.Stores, errChan chan error) {
	txmessage := &message.EditContactMainProcessToRenderer{
		Record: rxmessage.Record,
	}
	if err := stores.Contact.Update(rxmessage.Record); err != nil {
		errmsg := fmt.Sprintf("handleEditContact: stores.Contact.Update(rxmessage.Record): error is %s\n", err.Error())
		log.Println(errmsg)
		// Send back the error.
		txmessage.Error = true
		txmessage.ErrorMessage = errmsg
		sending <- txmessage
		return
	}
	// Send the record back with no error.
	sending <- txmessage
	// Also send the message that the contacts store has been modified.
	sending <- &message.ReloadContactsMainProcessToRenderer{}
	return
}
