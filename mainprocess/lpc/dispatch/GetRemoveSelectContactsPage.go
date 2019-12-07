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

// handleGetRemoveSelectContactsPage is the *message.GetRemoveSelectContactsPageRendererToMainProcess handler.
// It's response back to the renderer is the *message.GetRemoveSelectContactsPageMainProcessToRenderer.
// Param ctx is the context. if <-ctx.Done() then the main process is shutting down.
// Param rxmessage *message.GetRemoveSelectContactsPageRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.GetRemoveSelectContactsPageMainProcessToRenderer message back to the renderer.
// Param stores is a struct the contains each of your stores.
// Param errChan is the channel to send the handler's error through since the handler does not return it's error.
func handleGetRemoveSelectContactsPage(ctx context.Context, rxmessage *message.GetRemoveSelectContactsPageRendererToMainProcess, sending lpc.Sending, stores *store.Stores, errChan chan error) {
	txmessage := &message.GetRemoveSelectContactsPageMainProcessToRenderer{}
	var rr []*record.Contact
	var err error
	if rr, err = stores.Contact.GetAll(); err != nil {
		txmessage.Error = true
		txmessage.ErrorMessage = err.Error()
		sending <- txmessage
		return
	}
	sorted := sortRecords(rr)
	selected := selectRecords(sorted, rxmessage.FirstRecordIndex, rxmessage.Max)
	txmessage.FirstRecordIndex = rxmessage.FirstRecordIndex
	txmessage.State = rxmessage.State
	txmessage.Records = selected
	txmessage.TotalRecordCount = uint64(len(rr))
	sending <- txmessage
	return
}
