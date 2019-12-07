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

// handleReloadContacts is the *message.ReloadContactsRendererToMainProcess handler.
// It's response back to the renderer is the *message.ReloadContactsMainProcessToRenderer.
// Param ctx is the context. if <-ctx.Done() then the main process is shutting down.
// Param rxmessage *message.ReloadContactsRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.ReloadContactsMainProcessToRenderer message back to the renderer.
// Param stores is a struct the contains each of your stores.
// Param errChan is the channel to send the handler's error through since the handler does not return it's error.
func handleReloadContacts(ctx context.Context, rxmessage *message.ReloadContactsRendererToMainProcess, sending lpc.Sending, stores *store.Stores, errChan chan error) {
	return
}
