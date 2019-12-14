// +build js, wasm

package printselectpanel

import (
	"fmt"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/josephbudd/crud/rendererprocess/api/event"
	"github.com/josephbudd/crud/rendererprocess/kickwasmwidgets"
	"github.com/josephbudd/crud/rendererprocess/api/markup"
	"github.com/pkg/errors"
)

/*

	Panel name: PrintSelectPanel

*/

// panelController controls user input.
type panelController struct {
	group     *panelGroup
	presenter *panelPresenter
	messenger *panelMessenger

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelController fields.

	// example:

	import "syscall/js"
	import "github.com/josephbudd/crud/rendererprocess/api/markup"

	addCustomerName   *markup.Element
	addCustomerSubmit *markup.Element

	*/

	printSelectWidgetWrapper *markup.Element
	vlist                    *kickwasmwidgets.VList
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
	if controller.addCustomerName = document.ElementByID("addCustomerName"); controller.addCustomerName == nil {
		err = errors.New("unable to find #addCustomerName")
		return
	}

	// Define the submit button GUI controller.
	if controller.addCustomerSubmit = document.ElementByID("addCustomerSubmit"); controller.addCustomerSubmit == nil {
		err = errors.New("unable to find #addCustomerSubmit")
		return
	}
	// Handle the submit button's onclick event.
	controller.addCustomerSubmit.SetEventHandler(controller.handleSubmit, "click", false)

	*/

	// Define the vlist wrapper div and set it's handler.
	if controller.printSelectWidgetWrapper = document.ElementByID("printSelectWidgetWrapper"); controller.printSelectWidgetWrapper == nil {
		err = errors.New("unable to find #printSelectWidgetWrapper")
		return
	}
	// Define the vlist.
	// The vlist ID is it's own unique vlist state.
	vls := kickwasmwidgets.NewVListState()
	vlistID := vls.GetNextState()
	controller.vlist = kickwasmwidgets.NewVList(controller.printSelectWidgetWrapper,
		vlistID,
		// max real contacts in the virtual list of all contacts.
		50,
		// On size shows the select panel.
		func() {
			controller.group.showPrintSelectPanel(false)
		},
		// On no size shows the not ready panel.
		func() {
			controller.group.showPrintNotReadyPanel(false)
		},
		// needToInitializeFunc func(count, state uint64),
		func(count, state uint64) {
			controller.messenger.getPrintSelectContactsPage(0, count, state)
		},
		// needToPrependFunc func(button *markup.Element, count, state uint64)
		// Param button is the current first button.
		// Param count is the number of buttons that the selector needs.
		// Param state is the selector's ID OR'd with it's current behavior ( prepending ).
		func(button *markup.Element, count, state uint64) {
			// Remember. Each button is a contact record.
			// Find the records that will preceed the current button's record.
			buttonIndex, _ := button.AttributeUint64(sortedIndexAttributeName)
			startIndex := buttonIndex - count
			if startIndex < 0 {
				startIndex = 0
			}
			actualCount := 1 + (buttonIndex - startIndex)
			controller.messenger.getPrintSelectContactsPage(startIndex, actualCount, state)
		},
		// needToAppendFunc func(button *markup.Element, count, state uint64)
		// Param button is the current last button.
		// Param count is the number of buttons that the selector needs.
		// Param state is the selector's ID OR'd with it's current behavior ( appending ).
		func(button *markup.Element, count, state uint64) {
			buttonIndex, _ := button.AttributeUint64(sortedIndexAttributeName)
			controller.messenger.getPrintSelectContactsPage(buttonIndex, count, state)
		},
		document,
	)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

// example:

import "github.com/josephbudd/crud/domain/store/record"
import "github.com/josephbudd/crud/rendererprocess/api/event"
import "github.com/josephbudd/crud/rendererprocess/api/display"

func (controller *panelController) handleSubmit(e event.Event) (nilReturn interface{}) {
	// See renderer/event/event.go.
	// The event.Event funcs.
	//   e.PreventDefaultBehavior()
	//   e.StopCurrentPhasePropagation()
	//   e.StopAllPhasePropagation()
	//   target := e.JSTarget
	//   event := e.JSEvent
	// You must use the javascript event e.JSEvent, as a js.Value.
	// However, you can use the target as a *markup.Element
	//   target := document.NewElementFromJSValue(e.JSTarget)

	name := strings.TrimSpace(controller.addCustomerName.Value())
	if len(name) == 0 {
		display.Error("Customer Name is required.")
		return
	}
	r := &record.Customer{
		Name: name,
	}
	controller.messenger.AddCustomer(r)
	return
}

*/

// buildButtons is called by the messenger when the main process sends back the message reponse to the vilsts request for more records.
// buildButtons converts the records to buttons for the vlist.
// Then the new buttons, state and total record count, are passed to the vlist's Build func.
// vlist.Build rebuilds the vlist according to state which will be initialize, prepend, or append.
// Each button also has it's own onclick handler
//   which calls the messenger's getContact func so that a contact record can be printed.
func (controller *panelController) buildButtons(rr []*record.Contact, startSortedIndex, totalRecordCount, state uint64) {
	// convert the records into HTML buttons for the vlist
	buttons := make([]*markup.Element, len(rr))
	for i, r := range rr {
		button := document.NewBUTTON()
		// Button attributes.
		// Sorted index is needed for prepending and appending more buttons to the vlist.
		button.SetAttribute(sortedIndexAttributeName, startSortedIndex+uint64(i))
		// Record id needed when the user selects the record to print ( onclick ).
		button.SetAttribute(recordIDAttributeName, r.ID)
		// Button onclick
		f := func(e event.Event) interface{} {
			// the user has clicked on a button
			target := document.NewElementFromJSValue(e.JSTarget)
			for {
				if target.TagName() == "BUTTON" {
					break
				}
				target = target.Parent()
			}
			id, _ := target.AttributeUint64(recordIDAttributeName)
			controller.messenger.getContact(id)
			return nil
		}
		button.SetEventHandler(f, "click", false)
		// Button label.
		// Name
		h4 := document.NewH4()
		h4.SetInnerText(r.Name)
		button.AppendChild(h4)
		// Address 1 & 2
		p := document.NewP()
		p.AppendText(r.Address1)
		if len(r.Address2) > 0 {
			p.AppendChild(document.NewBR())
			p.AppendText(r.Address2)
		}
		// City, state zip.
		p.AppendChild(document.NewBR())
		p.AppendText(fmt.Sprintf("%s, %s %s", r.City, r.State, r.Zip))
		button.AppendChild(p)
		// Email, phone, social.
		p = document.NewP()
		p.AppendChild(document.NewBR())
		p.AppendText(r.Email)
		p.AppendChild(document.NewBR())
		p.AppendText(r.Phone)
		p.AppendChild(document.NewBR())
		p.AppendText(r.Social)
		button.AppendChild(p)
		buttons[i] = button
	}
	controller.vlist.Build(buttons, state, totalRecordCount)
}

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	// example:

	controller.customerSelectWidget.start()

	*/

	controller.vlist.Start()
}
