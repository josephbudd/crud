package addpanel

import (
	"syscall/js"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/pkg/errors"
)

/*

	Panel name: AddPanel

*/

// panelController controls user input.
type panelController struct {
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller

	/* NOTE TO DEVELOPER. Step 1 of 4.

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
	cb := tools.RegisterEventCallBack(controller.handleSubmit, true, true, true)
	notJS.SetOnClick(controller.contactAddSubmit, cb)

	if controller.contactAddCancel = notJS.GetElementByID("contactAddCancel"); controller.contactAddCancel == null {
		err = errors.New("unable to find #contactAddCancel")
		return
	}
	cb = tools.RegisterEventCallBack(controller.handleCancel, true, true, true)
	notJS.SetOnClick(controller.contactAddCancel, cb)

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
	controller.caller.addContact(r)
	return
}

func (controller *panelController) handleCancel(event js.Value) (nilReturn interface{}) {
	controller.presenter.clearForm()
	tools.Back()
	return
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

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/

}
