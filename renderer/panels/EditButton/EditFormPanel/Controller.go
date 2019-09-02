package editformpanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/josephbudd/crud/renderer/viewtools"
)

/*

	Panel name: EditFormPanel

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

	contactEditName     js.Value
	contactEditAddress1 js.Value
	contactEditAddress2 js.Value
	contactEditCity     js.Value
	contactEditState    js.Value
	contactEditZip      js.Value
	contactEditPhone    js.Value
	contactEditEmail    js.Value
	contactEditSocial   js.Value
	contactEditSubmit   js.Value
	contactEditCancel   js.Value

	contact *record.Contact
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

	if controller.contactEditName = notJS.GetElementByID("contactEditName"); controller.contactEditName == null {
		err = errors.New("unable to find #contactEditName")
		return
	}
	if controller.contactEditAddress1 = notJS.GetElementByID("contactEditAddress1"); controller.contactEditAddress1 == null {
		err = errors.New("unable to find #contactEditAddress1")
		return
	}
	if controller.contactEditAddress2 = notJS.GetElementByID("contactEditAddress2"); controller.contactEditAddress2 == null {
		err = errors.New("unable to find #contactEditAddress2")
		return
	}
	if controller.contactEditCity = notJS.GetElementByID("contactEditCity"); controller.contactEditCity == null {
		err = errors.New("unable to find #contactEditCity")
		return
	}
	if controller.contactEditState = notJS.GetElementByID("contactEditState"); controller.contactEditState == null {
		err = errors.New("unable to find #contactEditState")
		return
	}
	if controller.contactEditZip = notJS.GetElementByID("contactEditZip"); controller.contactEditZip == null {
		err = errors.New("unable to find #contactEditZip")
		return
	}
	if controller.contactEditPhone = notJS.GetElementByID("contactEditPhone"); controller.contactEditPhone == null {
		err = errors.New("unable to find #contactEditPhone")
		return
	}
	if controller.contactEditEmail = notJS.GetElementByID("contactEditEmail"); controller.contactEditEmail == null {
		err = errors.New("unable to find #contactEditEmail")
		return
	}
	if controller.contactEditSocial = notJS.GetElementByID("contactEditSocial"); controller.contactEditSocial == null {
		err = errors.New("unable to find #contactEditSocial")
		return
	}

	if controller.contactEditSubmit = notJS.GetElementByID("contactEditSubmit"); controller.contactEditSubmit == null {
		err = errors.New("unable to find #contactEditSubmit")
		return
	}
	// Receive the submit button's onclick event.
	controller.receiveEvent(controller.contactEditSubmit, "onclick", false, false, false)

	if controller.contactEditCancel = notJS.GetElementByID("contactEditCancel"); controller.contactEditCancel == null {
		err = errors.New("unable to find #contactEditCancel")
		return
	}
	// Receive the cancel button's onclick event.
	controller.receiveEvent(controller.contactEditCancel, "onclick", false, false, false)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 5.

// Handlers and other functions.

*/

func (controller *panelController) handleSubmit(event js.Value) {
	r := controller.getRecord()
	if len(r.Name) == 0 {
		tools.Error("Name is required.")
		return
	}
	if len(r.Address1) == 0 {
		tools.Error("Address1 is required.")
		return
	}
	if len(r.City) == 0 {
		tools.Error("City is required.")
		return
	}
	if len(r.State) == 0 {
		tools.Error("State is required.")
		return
	}
	if len(r.Zip) == 0 {
		tools.Error("Zip is required.")
		return
	}
	if len(r.Email) == 0 && len(r.Phone) == 0 {
		tools.Error("Either Email or Phone is required.")
		return
	}
	controller.caller.editContact(r)
}

func (controller *panelController) handleCancel(event js.Value) {
	controller.presenter.clearForm()
	// Go back to the EditSelectPanel
	controller.group.showEditSelectPanel(false)
}

func (controller *panelController) handleGetContact(r *record.Contact) {
	controller.contact = r
	controller.presenter.fillForm(r)
	controller.group.showEditFormPanel(false)
}

func (controller *panelController) getRecord() (r *record.Contact) {
	r = &record.Contact{
		ID:       controller.contact.ID,
		Name:     notJS.GetValue(controller.contactEditName),
		Address1: notJS.GetValue(controller.contactEditAddress1),
		Address2: notJS.GetValue(controller.contactEditAddress2),
		City:     notJS.GetValue(controller.contactEditCity),
		State:    notJS.GetValue(controller.contactEditState),
		Zip:      notJS.GetValue(controller.contactEditZip),
		Phone:    notJS.GetValue(controller.contactEditPhone),
		Email:    notJS.GetValue(controller.contactEditEmail),
		Social:   notJS.GetValue(controller.contactEditSocial),
	}
	return
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

				case controller.contactEditSubmit:
					controller.handleSubmit(event.Event)
				case controller.contactEditCancel:
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
