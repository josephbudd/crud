package removeformpanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/josephbudd/crud/renderer/viewtools"
)

/*

	Panel name: RemoveFormPanel

*/

// panelController controls user input.
type panelController struct {
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller
	eventCh   chan viewtools.Event

	/* NOTE TO DEVELOPER. Step 1 of 5.

	// Declare your panelController members.

	*/

	contact             *record.Contact
	contactRemoveSubmit js.Value
	contactRemoveCancel js.Value
}

// defineControlsReceiveEvents defines controller members and starts receiving their events.
// Returns the error.
func (controller *panelController) defineControlsReceiveEvents() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(controller *panelController) defineControlsReceiveEvents()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 5.

	// Define the controller members by their html elements.
	// Receive their events.
	// example:

	// Define the customer name input field.
	if controller.addCustomerName = notJS.GetElementByID("addCustomerName"); controller.addCustomerName == null {
		err = errors.New("unable to find #addCustomerName")
		return
	}

	// Define the submit button.
	if controller.addCustomerSubmit = notJS.GetElementByID("addCustomerSubmit"); controller.addCustomerSubmit == null {
		err = errors.New("unable to find #addCustomerSubmit")
		return
	}
	// Receive the submit button's onclick event.
	controller.receiveEvent(controller.addCustomerSubmit, "onclick", false, false, false)

	*/

	if controller.contactRemoveSubmit = notJS.GetElementByID("contactRemoveSubmit"); controller.contactRemoveSubmit == null {
		err = errors.New("unable to find #contactRemoveSubmit")
		return
	}
	// Receive the remove button's onclick event.
	controller.receiveEvent(controller.contactRemoveSubmit, "onclick", false, false, false)

	if controller.contactRemoveCancel = notJS.GetElementByID("contactRemoveCancel"); controller.contactRemoveCancel == null {
		err = errors.New("unable to find #contactRemoveCancel")
		return
	}
	// Receive the cancel button's onclick event.
	controller.receiveEvent(controller.contactRemoveCancel, "onclick", false, false, false)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 5.

// Handlers and other functions.

*/

func (controller *panelController) handleSubmit(event js.Value) {
	controller.caller.removeContact(controller.contact.ID)
}

func (controller *panelController) handleCancel(event js.Value) {
	controller.presenter.clearForm()
	// Go back to the RemoveSelectPanel
	controller.group.showRemoveSelectPanel(false)
}

func (controller *panelController) handleGetContact(r *record.Contact) {
	controller.contact = r
	controller.presenter.fillForm(r)
	controller.group.showRemoveFormPanel(false)
}

// dispatchEvents dispatches events from the controls.
// It stops when it receives on the eoj channel.
func (controller *panelController) dispatchEvents() {
	go func() {
		var event viewtools.Event
		for {
			select {
			case <-eojCh:
				return
			case event = <-controller.eventCh:
				// An event that this controller is receiving from one of its members.
				switch event.Target {

				/* NOTE TO DEVELOPER. Step 4 of 5.

				// 4.1.a: Add a case for each controller member
				//          that you are receiving events for.
				// 4.1.b: In that case statement, pass the event to your event handler.

				*/

				case controller.contactRemoveSubmit:
					controller.handleSubmit(event.Event)
				case controller.contactRemoveCancel:
					controller.handleCancel(event.Event)
				}
			}
		}
	}()

	return
}

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 5 of 5.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/
}

// receiveEvent gets this controller listening for element's event.
// Param elements if the controler's element.
// Param event is the event ex: "onclick".
// Param preventDefault indicates if the default behavior of the event must be prevented.
// Param stopPropagation indicates if the event's propogation must be stopped.
// Param stopImmediatePropagation indicates if the event's immediate propogation must be stopped.
func (controller *panelController) receiveEvent(element js.Value, event string, preventDefault, stopPropagation, stopImmediatePropagation bool) {
	tools.SendEvent(controller.eventCh, element, event, preventDefault, stopPropagation, stopImmediatePropagation)
}
