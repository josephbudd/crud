package dispatch

import (
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
// Param rxmessage *message.RemoveContactRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.RemoveContactMainProcessToRenderer message back to the renderer.
// Param eojing lpc.EOJer ( End Of Job ) is an interface for your go routine to receive a stop signal.
//   It signals go routines that they must stop because the main process is ending.
//   So only use it inside a go routine if you have one.
//   In your go routine
//     1. Get a channel to listen to with eojing.NewEOJ().
//     2. Before your go routine returns, release that channel with eojing.Release().
// Param stores is a struct the contains each of your stores.
func handleRemoveContact(rxmessage *message.RemoveContactRendererToMainProcess, sending lpc.Sending, eojing lpc.EOJer, stores *store.Stores) {
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
