package editselectpanel

import (
	"fmt"
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/josephbudd/crud/renderer/kickwasmwidgets"
	"github.com/josephbudd/crud/renderer/viewtools"
)

/*

	Panel name: EditSelectPanel

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

	editSelectWidgetWrapper  js.Value
	vlist                    *kickwasmwidgets.VList
	sortedIndexAttributeName string
	recordIDAttributeName    string
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

	controller.sortedIndexAttributeName = "sortedIndex"
	controller.recordIDAttributeName = "recordID"

	// Define the vlist wrapper div and set it's handler.
	if controller.editSelectWidgetWrapper = notJS.GetElementByID("editSelectWidgetWrapper"); controller.editSelectWidgetWrapper == null {
		err = errors.New("unable to find #editSelectWidgetWrapper")
		return
	}
	// Define the vlist.
	// The vlist ID is it's own unique vlist state.
	vls := kickwasmwidgets.NewVListState()
	vlistID := vls.GetNextState()
	controller.vlist = kickwasmwidgets.NewVList(controller.editSelectWidgetWrapper,
		vlistID,
		// max real contacts in the virtual list of all contacts.
		50,
		// On size shows the select panel.
		func() {
			controller.group.showEditSelectPanel(false)
		},
		// On no size shows the not ready panel.
		func() {
			controller.group.showEditNotReadyPanel(false)
		},
		// needToInitializeFunc func(count, state uint64),
		func(count, state uint64) {
			controller.caller.getEditSelectContactsPage(0, count, state)
		},
		// needToPrependFunc func(button js.Value, count, state uint64)
		// Param button is the current first button.
		// Param count is the number of buttons that the selector needs.
		// Param state is the selector's ID OR'd with it's current behavior ( prepending ).
		func(button js.Value, count, state uint64) {
			// Remember. Each button is a contact record.
			// Find the records that will preceed the current button's record.
			buttonIndex := notJS.GetAttributeUint64(button, controller.sortedIndexAttributeName)
			startIndex := buttonIndex - count
			if startIndex < 0 {
				startIndex = 0
			}
			actualCount := 1 + (buttonIndex - startIndex)
			controller.caller.getEditSelectContactsPage(startIndex, actualCount, state)
		},
		// needToAppendFunc func(button js.Value, count, state uint64)
		// Param button is the current last button.
		// Param count is the number of buttons that the selector needs.
		// Param state is the selector's ID OR'd with it's current behavior ( appending ).
		func(button js.Value, count, state uint64) {
			buttonIndex := notJS.GetAttributeUint64(button, controller.sortedIndexAttributeName)
			controller.caller.getEditSelectContactsPage(buttonIndex, count, state)
		},
		notJS,
		tools,
	)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 5.

// Handlers and other functions.

*/

// buildButtons is called by the caller when the main process sends back the message reponse to the vilsts request for more records.
// buildButtons converts the records to buttons for the vlist.
// Then the new buttons, state and total record count, are passed to the vlist's Build func.
// vlist.Build rebuilds the vlist according to state which will be initialize, prepend, or append.
// Each button also has it's own onclick handler
//   which calls the caller's getContact func so that a contact record can be edited.
func (controller *panelController) buildButtons(rr []*record.Contact, startSortedIndex, totalRecordCount, state uint64) {
	// convert the records into HTML buttons for the vlist
	buttons := make([]js.Value, len(rr))
	for i, r := range rr {
		button := notJS.CreateElementBUTTON()
		// Button attributes.
		// Sorted index is needed for prepending and appending more buttons to the vlist.
		notJS.SetAttributeUint64(button, controller.sortedIndexAttributeName, startSortedIndex+uint64(i))
		// Record id needed when the user selects the record to edit ( onclick ).
		notJS.SetAttributeUint64(button, controller.recordIDAttributeName, r.ID)
		// Button onclick
		cb := tools.RegisterEventCallBack(
			func(event js.Value) interface{} {
				// the user has clicked on a button
				target := notJS.GetEventTarget(event)
				for {
					if notJS.TagName(target) == "BUTTON" {
						break
					}
					target = notJS.ParentNode(target)
				}
				id := notJS.GetAttributeUint64(target, controller.recordIDAttributeName)
				controller.caller.getContact(id)
				return nil
			},
			true, true, true,
		)
		notJS.SetOnClick(button, cb)
		// Button label.
		// Name
		h4 := notJS.CreateElementH4()
		tn := notJS.CreateTextNode(r.Name)
		notJS.AppendChild(h4, tn)
		notJS.AppendChild(button, h4)
		// Address 1 & 2
		p := notJS.CreateElementP()
		tn = notJS.CreateTextNode(r.Address1)
		notJS.AppendChild(p, tn)
		br := notJS.CreateElementBR()
		if len(r.Address2) > 0 {
			notJS.AppendChild(p, br)
			tn = notJS.CreateTextNode(r.Address2)
			notJS.AppendChild(p, tn)
		}
		// City, state zip.
		br = notJS.CreateElementBR()
		notJS.AppendChild(p, br)
		tn = notJS.CreateTextNode(fmt.Sprintf("%s, %s %s", r.City, r.State, r.Zip))
		notJS.AppendChild(p, tn)
		notJS.AppendChild(button, p)
		// Email, phone, social.
		p = notJS.CreateElementP()
		br = notJS.CreateElementBR()
		notJS.AppendChild(p, br)
		tn = notJS.CreateTextNode(r.Email)
		notJS.AppendChild(p, tn)
		br = notJS.CreateElementBR()
		notJS.AppendChild(p, br)
		tn = notJS.CreateTextNode(r.Phone)
		notJS.AppendChild(p, tn)
		br = notJS.CreateElementBR()
		notJS.AppendChild(p, br)
		tn = notJS.CreateTextNode(r.Social)
		notJS.AppendChild(p, tn)
		notJS.AppendChild(button, p)
		buttons[i] = button
	}
	controller.vlist.Build(buttons, state, totalRecordCount)
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

	controller.vlist.Start()
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
