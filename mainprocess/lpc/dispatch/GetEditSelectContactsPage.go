package dispatch

import (
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
// Param rxmessage *message.GetEditSelectContactsPageRendererToMainProcess is the message received from the renderer.
// Param sending is the channel to use to send a *message.GetEditSelectContactsPageMainProcessToRenderer message back to the renderer.
// Param eojing lpc.EOJer ( End Of Job ) is an interface for your go routine to receive a stop signal.
//   It signals go routines that they must stop because the main process is ending.
//   So only use it inside a go routine if you have one.
//   In your go routine
//     1. Get a channel to listen to with eojing.NewEOJ().
//     2. Before your go routine returns, release that channel with eojing.Release().
// Param stores is a struct the contains each of your stores.
func handleGetEditSelectContactsPage(rxmessage *message.GetEditSelectContactsPageRendererToMainProcess, sending lpc.Sending, eojing lpc.EOJer, stores *store.Stores) {
	txmessage := &message.GetEditSelectContactsPageMainProcessToRenderer{}
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
