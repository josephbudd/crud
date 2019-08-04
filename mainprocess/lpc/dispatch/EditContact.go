package dispatch

import (
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
// Param rxmessage *message.EditContactRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.EditContactMainProcessToRenderer message back to the renderer.
// Param eojing lpc.EOJer ( End Of Job ) is an interface for your go routine to receive a stop signal.
//   It signals go routines that they must stop because the main process is ending.
//   So only use it inside a go routine if you have one.
//   In your go routine
//     1. Get a channel to listen to with eojing.NewEOJ().
//     2. Before your go routine returns, release that channel with eojing.Release().
func handleEditContact(rxmessage *message.EditContactRendererToMainProcess, sending lpc.Sending, eojing lpc.EOJer, stores *store.Stores) {
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
