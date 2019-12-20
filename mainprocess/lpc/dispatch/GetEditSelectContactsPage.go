package dispatch

import (
	"context"
	"sort"

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

// handleGetEditSelectContactsPage is the *message.GetEditSelectContactsPageRendererToMainProcess handler.
// It's response back to the renderer is the *message.GetEditSelectContactsPageMainProcessToRenderer.
// Param ctx is the context. if <-ctx.Done() then the main process is shutting down.
// Param rxmessage *message.GetEditSelectContactsPageRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.GetEditSelectContactsPageMainProcessToRenderer message back to the renderer.
// Param stores is a struct the contains each of your stores.
// Param errChan is the channel to send the handler's error through since the handler does not return it's error.
func handleGetEditSelectContactsPage(ctx context.Context, rxmessage *message.GetEditSelectContactsPageRendererToMainProcess, sending lpc.Sending, stores *store.Stores, errChan chan error) {
	txmessage := &message.GetEditSelectContactsPageMainProcessToRenderer{}
	var rr []*record.Contact
	var err error
	if rr, err = stores.Contact.GetAll(); err != nil {
		// Send the err to package main.
		errChan <- err
		// Send the error to the renderer.
		// A bolt database error is fatal.
		txmessage.Fatal = true
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

func selectRecords(sorted []*record.Contact, firstRecordIndex, max uint64) (selected []*record.Contact) {
	var l uint64
	if l = uint64(len(sorted)); l == 0 {
		return
	}
	if firstRecordIndex > 0 && l <= firstRecordIndex+1 {
		// no more records to send
		return
	}
	start := firstRecordIndex
	end := start + max
	if end > l {
		end = l
	}
	selected = sorted[start:end]
	return
}

func sortRecords(rr []*record.Contact) (sorted []*record.Contact) {
	l := len(rr)
	tracker := make(map[string][]*record.Contact)
	keys := make([]string, 0, l)
	for _, r := range rr {
		if _, found := tracker[r.Name]; !found {
			tracker[r.Name] = make([]*record.Contact, 0, 10)
			keys = append(keys, r.Name)
		}
		tracker[r.Name] = append(tracker[r.Name], r)
	}
	sort.Strings(keys)
	sorted = make([]*record.Contact, 0, l)
	for _, name := range keys {
		for _, r := range tracker[name] {
			sorted = append(sorted, r)
		}
	}
	return
}
