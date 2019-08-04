package editselectpanel

import (
	"fmt"
	"syscall/js"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/josephbudd/crud/renderer/kickwasmwidgets"
	"github.com/pkg/errors"
)

/*

	Panel name: EditSelectPanel

*/

// panelControler controls user input.
type panelControler struct {
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelControler members.

	*/

	editSelectWidgetWrapper  js.Value
	vlist                    *kickwasmwidgets.VList
	sortedIndexAttributeName string
	recordIDAttributeName    string
}

// defineControlsSetHandlers defines controler members and sets their handlers.
// Returns the error.
func (controler *panelControler) defineControlsSetHandlers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(controler *panelControler) defineControlsSetHandlers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controler members by their html elements.
	// Set their handlers.

	*/

	controler.sortedIndexAttributeName = "sortedIndex"
	controler.recordIDAttributeName = "recordID"

	// Define the vlist wrapper div and set it's handler.
	if controler.editSelectWidgetWrapper = notJS.GetElementByID("editSelectWidgetWrapper"); controler.editSelectWidgetWrapper == null {
		err = errors.New("unable to find #editSelectWidgetWrapper")
		return
	}
	// Define the vlist.
	// The vlist ID is it's own unique vlist state.
	vls := kickwasmwidgets.NewVListState()
	vlistID := vls.GetNextState()
	controler.vlist = kickwasmwidgets.NewVList(controler.editSelectWidgetWrapper,
		vlistID,
		// max real contacts in the virtual list of all contacts.
		50,
		// On size shows the select panel.
		func() {
			controler.group.showEditSelectPanel(false)
		},
		// On no size shows the not ready panel.
		func() {
			controler.group.showEditNotReadyPanel(false)
		},
		// needToInitializeFunc func(count, state uint64),
		func(count, state uint64) {
			controler.caller.getEditSelectContactsPage(0, count, state)
		},
		// needToPrependFunc func(button js.Value, count, state uint64)
		// Param button is the current first button.
		// Param count is the number of buttons that the selector needs.
		// Param state is the selector's ID OR'd with it's current behavior ( prepending ).
		func(button js.Value, count, state uint64) {
			// Remember. Each button is a contact record.
			// Find the records that will preceed the current button's record.
			buttonIndex := notJS.GetAttributeUint64(button, controler.sortedIndexAttributeName)
			startIndex := buttonIndex - count
			if startIndex < 0 {
				startIndex = 0
			}
			actualCount := 1 + (buttonIndex - startIndex)
			controler.caller.getEditSelectContactsPage(startIndex, actualCount, state)
		},
		// needToAppendFunc func(button js.Value, count, state uint64)
		// Param button is the current last button.
		// Param count is the number of buttons that the selector needs.
		// Param state is the selector's ID OR'd with it's current behavior ( appending ).
		func(button js.Value, count, state uint64) {
			buttonIndex := notJS.GetAttributeUint64(button, controler.sortedIndexAttributeName)
			controler.caller.getEditSelectContactsPage(buttonIndex, count, state)
		},
		notJS,
		tools,
	)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

// buildButtons is called by the caller when the main process sends back the message reponse to the vilsts request for more records.
// buildButtons converts the records to buttons for the vlist.
// Then the new buttons, state and total record count, are passed to the vlist's Build func.
// vlist.Build rebuilds the vlist according to state which will be initialize, prepend, or append.
// Each button also has it's own onclick handler
//   which calls the caller's getContact func so that a contact record can be edited.
func (controler *panelControler) buildButtons(rr []*record.Contact, startSortedIndex, totalRecordCount, state uint64) {
	// convert the records into HTML buttons for the vlist
	buttons := make([]js.Value, len(rr))
	for i, r := range rr {
		button := notJS.CreateElementBUTTON()
		// Button attributes.
		// Sorted index is needed for prepending and appending more buttons to the vlist.
		notJS.SetAttributeUint64(button, controler.sortedIndexAttributeName, startSortedIndex+uint64(i))
		// Record id needed when the user selects the record to edit ( onclick ).
		notJS.SetAttributeUint64(button, controler.recordIDAttributeName, r.ID)
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
				id := notJS.GetAttributeUint64(target, controler.recordIDAttributeName)
				controler.caller.getContact(id)
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
	controler.vlist.Build(buttons, state, totalRecordCount)
	return
}

// initialCalls runs the first code that the controler needs to run.
func (controler *panelControler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/

	controler.vlist.Start()

}
