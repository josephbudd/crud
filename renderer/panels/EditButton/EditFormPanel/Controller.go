package editformpanel

import (
	"syscall/js"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/pkg/errors"
)

/*

	Panel name: EditFormPanel

*/

// panelController controls user input.
type panelController struct {
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller

	/* NOTE TO DEVELOPER. Step 1 of 4.

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

// defineControlsSetHandlers defines controller members and sets their handlers.
// Returns the error.
func (controller *panelController) defineControlsSetHandlers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(controller *panelController) defineControlsSetHandlers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define the Controller members by their html elements.
	// Set their handlers.

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
	cb := tools.RegisterEventCallBack(controller.handleSubmit, true, true, true)
	notJS.SetOnClick(controller.contactEditSubmit, cb)

	if controller.contactEditCancel = notJS.GetElementByID("contactEditCancel"); controller.contactEditCancel == null {
		err = errors.New("unable to find #contactEditCancel")
		return
	}
	cb = tools.RegisterEventCallBack(controller.handleCancel, true, true, true)
	notJS.SetOnClick(controller.contactEditCancel, cb)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (controller *panelController) handleSubmit(event js.Value) (nilReturn interface{}) {
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
	return
}

func (controller *panelController) handleCancel(event js.Value) (nilReturn interface{}) {
	controller.presenter.clearForm()
	// Go back to the EditSelectPanel
	controller.group.showEditSelectPanel(false)
	return
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

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/

}
