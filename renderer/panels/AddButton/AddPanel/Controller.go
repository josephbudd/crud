package addpanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/josephbudd/crud/renderer/viewtools"
)

/*

	Panel name: AddPanel

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

	contactAddName     js.Value
	contactAddAddress1 js.Value
	contactAddAddress2 js.Value
	contactAddCity     js.Value
	contactAddState    js.Value
	contactAddZip      js.Value
	contactAddPhone    js.Value
	contactAddEmail    js.Value
	contactAddSocial   js.Value
	contactAddSubmit   js.Value
	contactAddCancel   js.Value
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

	*/

	if controller.contactAddName = notJS.GetElementByID("contactAddName"); controller.contactAddName == null {
		err = errors.New("unable to find #contactAddName")
		return
	}
	if controller.contactAddAddress1 = notJS.GetElementByID("contactAddAddress1"); controller.contactAddAddress1 == null {
		err = errors.New("unable to find #contactAddAddress1")
		return
	}
	if controller.contactAddAddress2 = notJS.GetElementByID("contactAddAddress2"); controller.contactAddAddress2 == null {
		err = errors.New("unable to find #contactAddAddress2")
		return
	}
	if controller.contactAddCity = notJS.GetElementByID("contactAddCity"); controller.contactAddCity == null {
		err = errors.New("unable to find #contactAddCity")
		return
	}
	if controller.contactAddState = notJS.GetElementByID("contactAddState"); controller.contactAddState == null {
		err = errors.New("unable to find #contactAddState")
		return
	}
	if controller.contactAddZip = notJS.GetElementByID("contactAddZip"); controller.contactAddZip == null {
		err = errors.New("unable to find #contactAddZip")
		return
	}
	if controller.contactAddPhone = notJS.GetElementByID("contactAddPhone"); controller.contactAddPhone == null {
		err = errors.New("unable to find #contactAddPhone")
		return
	}
	if controller.contactAddEmail = notJS.GetElementByID("contactAddEmail"); controller.contactAddEmail == null {
		err = errors.New("unable to find #contactAddEmail")
		return
	}
	if controller.contactAddSocial = notJS.GetElementByID("contactAddSocial"); controller.contactAddSocial == null {
		err = errors.New("unable to find #contactAddSocial")
		return
	}

	if controller.contactAddSubmit = notJS.GetElementByID("contactAddSubmit"); controller.contactAddSubmit == null {
		err = errors.New("unable to find #contactAddSubmit")
		return
	}
	// Receive the add button's onclick event.
	controller.receiveEvent(controller.contactAddSubmit, "onclick", false, false, false)

	if controller.contactAddCancel = notJS.GetElementByID("contactAddCancel"); controller.contactAddCancel == null {
		err = errors.New("unable to find #contactAddCancel")
		return
	}
	// Receive the cancel button's onclick event.
	controller.receiveEvent(controller.contactAddCancel, "onclick", false, false, false)

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
	controller.caller.addContact(r)
}

func (controller *panelController) handleCancel(event js.Value) {
	controller.presenter.clearForm()
	tools.Back()
}

func (controller *panelController) getRecord() *record.Contact {
	return &record.Contact{
		Name:     notJS.GetValue(controller.contactAddName),
		Address1: notJS.GetValue(controller.contactAddAddress1),
		Address2: notJS.GetValue(controller.contactAddAddress2),
		City:     notJS.GetValue(controller.contactAddCity),
		State:    notJS.GetValue(controller.contactAddState),
		Zip:      notJS.GetValue(controller.contactAddZip),
		Phone:    notJS.GetValue(controller.contactAddPhone),
		Email:    notJS.GetValue(controller.contactAddEmail),
		Social:   notJS.GetValue(controller.contactAddSocial),
	}
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

				case controller.contactAddSubmit:
					controller.handleSubmit(event.Event)

				case controller.contactAddCancel:
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
