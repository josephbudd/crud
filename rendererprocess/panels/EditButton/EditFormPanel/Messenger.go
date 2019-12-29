// +build js, wasm

package editformpanel

import (
	"github.com/josephbudd/crud/domain/lpc/message"
	"github.com/josephbudd/crud/domain/store/record"
	"github.com/josephbudd/crud/rendererprocess/api/display"
)

/*

	Panel name: EditFormPanel

*/

// panelMessenger communicates with the main process via an asynchrounous connection.
type panelMessenger struct {
	group      *panelGroup
	presenter  *panelPresenter
	controller *panelController

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// 1.1: Declare your panelMessenger members.

	// example:

	state uint64

	*/
}

/* NOTE TO DEVELOPER. Step 2 of 4.

// 2.1: Define your funcs which send a message to the main process.
// 2.2: Define your funcs which receive a message from the main process.

// example:

import "github.com/josephbudd/crud/domain/store/record"
import "github.com/josephbudd/crud/domain/lpc/message"
import "github.com/josephbudd/crud/rendererprocess/api/display"

// Add Customer.

func (messenger *panelMessenger) addCustomer(r *record.Customer) {
	msg := &message.AddCustomerRendererToMainProcess{
		UniqueID: messenger.uniqueID,
		Record:   record,
	}
	sendCh <- msg
}

func (messenger *panelMessenger) addCustomerRX(msg *message.AddCustomerMainProcessToRenderer) {
	if msg.UniqueID == messenger.uniqueID {
		if msg.Error {
			display.Error(msg.ErrorMessage)
			return
		}
		// no errors
		display.Success("Customer Added.")
	}
}

*/

// Get Contact

func (messenger *panelMessenger) getContactRX(msg *message.GetEditContactMainProcessToRenderer) {
	if msg.Error {
		display.Error(msg.ErrorMessage)
		return
	}
	// no errors
	messenger.controller.handleGetContact(msg.Record)
}

// Edit Contact.

func (messenger *panelMessenger) editContact(r *record.Contact) {
	msg := &message.EditContactRendererToMainProcess{
		Record: r,
	}
	sendCh <- msg
}

func (messenger *panelMessenger) editContactRX(msg *message.EditContactMainProcessToRenderer) {
	if msg.Error {
		display.Error(msg.ErrorMessage)
		return
	}
	// no errors
	messenger.presenter.clearForm()
	display.Success("Contact Edited.")
}

// dispatchMessages dispatches LPC messages from the main process.
// It stops when context is done.
func (messenger *panelMessenger) dispatchMessages() {
	go func() {
		for {
			select {
			case <-rendererProcessCtx.Done():
				return
			case msg := <-receiveCh:
				// A message sent from the main process to the renderer.
				switch msg := msg.(type) {

				/* NOTE TO DEVELOPER. Step 3 of 4.

				// 3.1:   Remove the default clause below.
				// 3.2.a: Add a case for each of the messages
				//          that you are expecting from the main process.
				// 3.2.b: In that case statement, pass the message to your message receiver func.

				// example:

				import "github.com/josephbudd/crud/domain/lpc/message"

				case *message.AddCustomerMainProcessToRenderer:
					messenger.addCustomerRX(msg)

				*/

				case *message.GetEditContactMainProcessToRenderer:
					messenger.getContactRX(msg)
				case *message.EditContactMainProcessToRenderer:
					messenger.editContactRX(msg)
				}
			}
		}
	}()

	return
}

// initialSends sends the first messages to the main process.
func (messenger *panelMessenger) initialSends() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	//4.1: Send messages to the main process right when the app starts.

	// example:

	import "github.com/josephbudd/crud/domain/data/loglevels"
	import "github.com/josephbudd/crud/domain/lpc/message"

	msg := &message.LogRendererToMainProcess{
		Level:   loglevels.LogLevelInfo,
		Message: "Started",
	}
	sendCh <- msg

	*/
}
