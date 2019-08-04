package addpanel

import (
	"syscall/js"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/pkg/errors"
)

/*

	Panel name: AddPanel

*/

// panelControler controls user input.
type panelControler struct {
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelControler members.

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

	if controler.contactAddName = notJS.GetElementByID("contactAddName"); controler.contactAddName == null {
		err = errors.New("unable to find #contactAddName")
		return
	}
	if controler.contactAddAddress1 = notJS.GetElementByID("contactAddAddress1"); controler.contactAddAddress1 == null {
		err = errors.New("unable to find #contactAddAddress1")
		return
	}
	if controler.contactAddAddress2 = notJS.GetElementByID("contactAddAddress2"); controler.contactAddAddress2 == null {
		err = errors.New("unable to find #contactAddAddress2")
		return
	}
	if controler.contactAddCity = notJS.GetElementByID("contactAddCity"); controler.contactAddCity == null {
		err = errors.New("unable to find #contactAddCity")
		return
	}
	if controler.contactAddState = notJS.GetElementByID("contactAddState"); controler.contactAddState == null {
		err = errors.New("unable to find #contactAddState")
		return
	}
	if controler.contactAddZip = notJS.GetElementByID("contactAddZip"); controler.contactAddZip == null {
		err = errors.New("unable to find #contactAddZip")
		return
	}
	if controler.contactAddPhone = notJS.GetElementByID("contactAddPhone"); controler.contactAddPhone == null {
		err = errors.New("unable to find #contactAddPhone")
		return
	}
	if controler.contactAddEmail = notJS.GetElementByID("contactAddEmail"); controler.contactAddEmail == null {
		err = errors.New("unable to find #contactAddEmail")
		return
	}
	if controler.contactAddSocial = notJS.GetElementByID("contactAddSocial"); controler.contactAddSocial == null {
		err = errors.New("unable to find #contactAddSocial")
		return
	}

	if controler.contactAddSubmit = notJS.GetElementByID("contactAddSubmit"); controler.contactAddSubmit == null {
		err = errors.New("unable to find #contactAddSubmit")
		return
	}
	cb := tools.RegisterEventCallBack(controler.handleSubmit, true, true, true)
	notJS.SetOnClick(controler.contactAddSubmit, cb)

	if controler.contactAddCancel = notJS.GetElementByID("contactAddCancel"); controler.contactAddCancel == null {
		err = errors.New("unable to find #contactAddCancel")
		return
	}
	cb = tools.RegisterEventCallBack(controler.handleCancel, true, true, true)
	notJS.SetOnClick(controler.contactAddCancel, cb)

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
	controler.caller.addContact(r)
	return
}

func (controler *panelControler) handleCancel(event js.Value) (nilReturn interface{}) {
	controler.presenter.clearForm()
	tools.Back()
	return
}

func (controler *panelControler) getRecord() *record.Contact {
	return &record.Contact{
		Name:     notJS.GetValue(controler.contactAddName),
		Address1: notJS.GetValue(controler.contactAddAddress1),
		Address2: notJS.GetValue(controler.contactAddAddress2),
		City:     notJS.GetValue(controler.contactAddCity),
		State:    notJS.GetValue(controler.contactAddState),
		Zip:      notJS.GetValue(controler.contactAddZip),
		Phone:    notJS.GetValue(controler.contactAddPhone),
		Email:    notJS.GetValue(controler.contactAddEmail),
		Social:   notJS.GetValue(controler.contactAddSocial),
	}
}

// initialCalls runs the first code that the controler needs to run.
func (controler *panelControler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/

}
