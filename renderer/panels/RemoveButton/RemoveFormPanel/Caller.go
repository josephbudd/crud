package removeformpanel

import "github.com/josephbudd/crud/domain/lpc/message"

/*

	Panel name: RemoveFormPanel

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

// GetContact

// Receive the contact record.
func (caller *panelCaller) getContactRX(msg *message.GetRemoveContactMainProcessToRenderer) {
	if msg.Error {
		return
	}
	caller.controller.handleGetContact(msg.Record)
}

// RemoveContact

func (caller *panelCaller) removeContact(id uint64) {
	txmessage := &message.RemoveContactRendererToMainProcess{
		ID: id,
	}
	sendCh <- txmessage
}

func (caller *panelCaller) removeContactRX(msg *message.RemoveContactMainProcessToRenderer) {
	if msg.Error {
		tools.Error(msg.ErrorMessage)
		return
	}
	tools.Success("Contact Removed.")
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

				case *message.GetRemoveContactMainProcessToRenderer:
					caller.getContactRX(msg)

				case *message.RemoveContactMainProcessToRenderer:
					caller.removeContactRX(msg)
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

	// example:

	// import "github.com/josephbudd/crud/domain/data/loglevels"
	// import "github.com/josephbudd/crud/domain/lpc/message"

	msg := &message.LogRendererToMainProcess{
		Level:   loglevels.LogLevelInfo,
		Message: "Started",
	}
	sendCh <- msg

	*/
}
