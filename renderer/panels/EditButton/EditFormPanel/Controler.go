package editformpanel

import (
	"syscall/js"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/pkg/errors"
)

/*

	Panel name: EditFormPanel

*/

// panelControler controls user input.
type panelControler struct {
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelControler members.

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

	if controler.contactEditName = notJS.GetElementByID("contactEditName"); controler.contactEditName == null {
		err = errors.New("unable to find #contactEditName")
		return
	}
	if controler.contactEditAddress1 = notJS.GetElementByID("contactEditAddress1"); controler.contactEditAddress1 == null {
		err = errors.New("unable to find #contactEditAddress1")
		return
	}
	if controler.contactEditAddress2 = notJS.GetElementByID("contactEditAddress2"); controler.contactEditAddress2 == null {
		err = errors.New("unable to find #contactEditAddress2")
		return
	}
	if controler.contactEditCity = notJS.GetElementByID("contactEditCity"); controler.contactEditCity == null {
		err = errors.New("unable to find #contactEditCity")
		return
	}
	if controler.contactEditState = notJS.GetElementByID("contactEditState"); controler.contactEditState == null {
		err = errors.New("unable to find #contactEditState")
		return
	}
	if controler.contactEditZip = notJS.GetElementByID("contactEditZip"); controler.contactEditZip == null {
		err = errors.New("unable to find #contactEditZip")
		return
	}
	if controler.contactEditPhone = notJS.GetElementByID("contactEditPhone"); controler.contactEditPhone == null {
		err = errors.New("unable to find #contactEditPhone")
		return
	}
	if controler.contactEditEmail = notJS.GetElementByID("contactEditEmail"); controler.contactEditEmail == null {
		err = errors.New("unable to find #contactEditEmail")
		return
	}
	if controler.contactEditSocial = notJS.GetElementByID("contactEditSocial"); controler.contactEditSocial == null {
		err = errors.New("unable to find #contactEditSocial")
		return
	}

	if controler.contactEditSubmit = notJS.GetElementByID("contactEditSubmit"); controler.contactEditSubmit == null {
		err = errors.New("unable to find #contactEditSubmit")
		return
	}
	cb := tools.RegisterEventCallBack(controler.handleSubmit, true, true, true)
	notJS.SetOnClick(controler.contactEditSubmit, cb)

	if controler.contactEditCancel = notJS.GetElementByID("contactEditCancel"); controler.contactEditCancel == null {
		err = errors.New("unable to find #contactEditCancel")
		return
	}
	cb = tools.RegisterEventCallBack(controler.handleCancel, true, true, true)
	notJS.SetOnClick(controler.contactEditCancel, cb)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (controler *panelControler) handleSubmit(event js.Value) (nilReturn interface{}) {
	r := controler.getRecord()
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
	controler.caller.editContact(r)
	return
}

func (controler *panelControler) handleCancel(event js.Value) (nilReturn interface{}) {
	controler.presenter.clearForm()
	// Go back to the EditSelectPanel
	controler.group.showEditSelectPanel(false)
	return
}

func (controler *panelControler) handleGetContact(r *record.Contact) {
	controler.contact = r
	controler.presenter.fillForm(r)
	controler.group.showEditFormPanel(false)
}

func (controler *panelControler) getRecord() (r *record.Contact) {
	r = &record.Contact{
		ID:       controler.contact.ID,
		Name:     notJS.GetValue(controler.contactEditName),
		Address1: notJS.GetValue(controler.contactEditAddress1),
		Address2: notJS.GetValue(controler.contactEditAddress2),
		City:     notJS.GetValue(controler.contactEditCity),
		State:    notJS.GetValue(controler.contactEditState),
		Zip:      notJS.GetValue(controler.contactEditZip),
		Phone:    notJS.GetValue(controler.contactEditPhone),
		Email:    notJS.GetValue(controler.contactEditEmail),
		Social:   notJS.GetValue(controler.contactEditSocial),
	}
	return
}

// initialCalls runs the first code that the controler needs to run.
func (controler *panelControler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/

}
