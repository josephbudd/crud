package removeselectpanel

import (
	"fmt"
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/josephbudd/crud/renderer/kickwasmwidgets"
	"github.com/josephbudd/crud/renderer/viewtools"
)

/*

	Panel name: RemoveSelectPanel

*/

// panelController controls user input.
type panelController struct {
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller
	eventCh   chan viewtools.Event

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelController fields.

	*/

	removeSelectWidgetWrapper js.Value
	vlist                     *kickwasmwidgets.VList
	sortedIndexAttributeName  string
	recordIDAttributeName     string
}

// defineControlsHandlers defines the GUI's controllers and their event handlers.
// Returns the error.
func (controller *panelController) defineControlsHandlers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(controller *panelController) defineControlsHandlers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define each controller in the GUI by it's html element.
	// Handle each controller's events.

	// example:

	// Define the customer name text input GUI controller.
	if controller.addCustomerName = notJS.GetElementByID("addCustomerName"); controller.addCustomerName == null {
		err = errors.New("unable to find #addCustomerName")
		return
	}

	// Define the submit button GUI controller.
	if controller.addCustomerSubmit = notJS.GetElementByID("addCustomerSubmit"); controller.addCustomerSubmit == null {
		err = errors.New("unable to find #addCustomerSubmit")
		return
	}
	// Handle the submit button's onclick event.
	tools.AddEventHandler(controller.handleSubmit, controller.addCustomerSubmit, "click", false)

	*/

	// Define the vlist wrapper div and set it's handler.
	if controller.removeSelectWidgetWrapper = notJS.GetElementByID("removeSelectWidgetWrapper"); controller.removeSelectWidgetWrapper == null {
		err = errors.New("unable to find #removeSelectWidgetWrapper")
		return
	}
	// Define the vlist.
	// The vlist ID is it's own unique vlist state.
	vls := kickwasmwidgets.NewVListState()
	vlistID := vls.GetNextState()
	controller.vlist = kickwasmwidgets.NewVList(controller.removeSelectWidgetWrapper,
		vlistID,
		// max real contacts in the virtual list of all contacts.
		50,
		// On size shows the select panel.
		func() {
			controller.group.showRemoveSelectPanel(false)
		},
		// On no size shows the not ready panel.
		func() {
			controller.group.showRemoveNotReadyPanel(false)
		},
		// needToInitializeFunc func(count, state uint64),
		func(count, state uint64) {
			controller.caller.getRemoveSelectContactsPage(0, count, state)
		},
		// needToPrependFunc func(button js.Value, count, state uint64)
		// Param button is the current first button.
		// Param count is the number of buttons that the selector needs.
		// Param state is the selector's ID OR'd with it's current behavior ( prepending ).
		func(button js.Value, count, state uint64) {
			// Remember. Each button is a contact record.
			// Find the records that will preceed the current button's record.
			buttonIndex := notJS.GetAttributeUint64(button, sortedIndexAttributeName)
			startIndex := buttonIndex - count
			if startIndex < 0 {
				startIndex = 0
			}
			actualCount := 1 + (buttonIndex - startIndex)
			controller.caller.getRemoveSelectContactsPage(startIndex, actualCount, state)
		},
		// needToAppendFunc func(button js.Value, count, state uint64)
		// Param button is the current last button.
		// Param count is the number of buttons that the selector needs.
		// Param state is the selector's ID OR'd with it's current behavior ( appending ).
		func(button js.Value, count, state uint64) {
			buttonIndex := notJS.GetAttributeUint64(button, sortedIndexAttributeName)
			controller.caller.getRemoveSelectContactsPage(buttonIndex, count, state)
		},
		notJS,
		tools,
	)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

// example:

// import "github.com/josephbudd/crud/domain/store/record"

func (controller *panelController) handleSubmit(e viewtools.Event) (nilReturn interface{}) {
	// See renderer/viewtools/event.go.
	// The viewtools.Event funcs.
	//   e.PreventDefaultBehavior()
	//   e.StopCurrentPhasePropagation()
	//   e.StopAllPhasePropagation()
	//   target := e.Target
	//   event := e.Event

	name := strings.TrimSpace(notJS.GetValue(controller.addCustomerName))
	if len(name) == 0 {
		tools.Error("Customer Name is required.")
		return
	}
	r := &record.Customer{
		Name: name,
	}
	controller.caller.AddCustomer(r)
	return
}

*/

// buildButtons is called by the caller when the main process sends back the message reponse to the vilsts request for more records.
// buildButtons converts the records to buttons for the vlist.
// Then the new buttons, state and total record count, are passed to the vlist's Build func.
// vlist.Build rebuilds the vlist according to state which will be initialize, prepend, or append.
// Each button also has it's own onclick handler
//   which calls the caller's getContact func so that a contact record can be confirmed or canceld for removal.
func (controller *panelController) buildButtons(rr []*record.Contact, startSortedIndex, totalRecordCount, state uint64) {
	// convert the records into HTML buttons for the vlist
	buttons := make([]js.Value, len(rr))
	for i, r := range rr {
		button := notJS.CreateElementBUTTON()
		// Button attributes.
		// Sorted index is needed for prepending and appending more buttons to the vlist.
		notJS.SetAttributeUint64(button, sortedIndexAttributeName, startSortedIndex+uint64(i))
		// Record id needed when the user selects the record to remove ( onclick ).
		notJS.SetAttributeUint64(button, recordIDAttributeName, r.ID)
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
				id := notJS.GetAttributeUint64(target, recordIDAttributeName)
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
	return
}

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/

	controller.vlist.Start()
}
