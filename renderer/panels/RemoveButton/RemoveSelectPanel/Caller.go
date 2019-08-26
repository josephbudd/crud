package removeselectpanel

import "github.com/josephbudd/crud/domain/lpc/message"

/*

	Panel name: RemoveSelectPanel

*/

// panelCaller communicates with the main process via an asynchrounous connection.
type panelCaller struct {
	group      *panelGroup
	presenter  *panelPresenter
	controller *panelController

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// 1.1: Declare your panelCaller members.

	*/
}

/* NOTE TO DEVELOPER. Step 2 of 4.

// 2.1: Define your funcs which send a message to the main process.
// 2.2: Define your funcs which receive a message from the main process.

*/

// Tell the main process to send back a page of records.
func (caller *panelCaller) getRemoveSelectContactsPage(firstRecordIndex, max, state uint64) {
	msg := &message.GetRemoveSelectContactsPageRendererToMainProcess{
		FirstRecordIndex: firstRecordIndex,
		Max:              max,
		State:            state,
	}
	sendCh <- msg
}

// Receive the page of records from the main process.
func (caller *panelCaller) getRemoveSelectContactsPageRX(msg *message.GetRemoveSelectContactsPageMainProcessToRenderer) {
	if msg.Error {
		tools.Error(msg.ErrorMessage)
		return
	}
	// No errors so pass the info to the controller.
	caller.controller.buildButtons(msg.Records, msg.FirstRecordIndex, msg.TotalRecordCount, msg.State)
}

// GetContact

// Tells the main process to send back a contact record for removal confirmation.
func (caller *panelCaller) getContact(id uint64) {
	msg := &message.GetRemoveContactRendererToMainProcess{
		ID: id,
	}
	sendCh <- msg
}

// Receive the contact record.
func (caller *panelCaller) getContactRX(msg *message.GetRemoveContactMainProcessToRenderer) {
	// This caller sent the message to the main process so this caller will handle the error if there is one.
	if msg.Error {
		tools.Error(msg.ErrorMessage)
		return
	}
	// No error so let the Remove panel handle this.
}

// Reload

func (caller *panelCaller) reloadContactsRX(msg *message.ReloadContactsMainProcessToRenderer) {
	// Tell the vlist to restart.
	caller.controller.vlist.Start()
}

// listen listens for messages from the main process.
// It stops when it receives on the eoj channel.
func (caller *panelCaller) listen() {
	go func() {
		for {
			select {
			case <-eojCh:
				return
			case msg := <-receiveCh:
				// A message sent from the main process to the renderer.
				switch msg := msg.(type) {

				/* NOTE TO DEVELOPER. Step 3 of 4.

				// 3.1:   Remove the default statement below.
				// 3.2.a: Add a case for each of the messages
				//          that you are expecting from the main process.
				// 3.2.b: In that case statement, pass the message to your message receiver func.

				*/

				case *message.GetRemoveSelectContactsPageMainProcessToRenderer:
					caller.getRemoveSelectContactsPageRX(msg)

				case *message.GetRemoveContactMainProcessToRenderer:
					caller.getContactRX(msg)

				case *message.ReloadContactsMainProcessToRenderer:
					caller.reloadContactsRX(msg)
				}
			}
		}
	}()

	return
}

// initialCalls makes the first calls to the main process.
func (caller *panelCaller) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	//4.1: Make any initial calls to the main process that must be made when the app starts.

	*/
}
