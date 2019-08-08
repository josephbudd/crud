package editformpanel

import (
	"syscall/js"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/pkg/errors"
)

/*

	Panel name: EditFormPanel

*/

// panelPresenter writes to the panel
type panelPresenter struct {
	group      *panelGroup
	controller *panelController
	caller     *panelCaller

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your panelPresenter members here.

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
}

// defineMembers defines the panelPresenter members by their html elements.
// Returns the error.
func (presenter *panelPresenter) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(presenter *panelPresenter) defineMembers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your panelPresenter members.

	*/

	if presenter.contactEditName = notJS.GetElementByID("contactEditName"); presenter.contactEditName == null {
		err = errors.New("unable to find #contactEditName")
		return
	}
	if presenter.contactEditAddress1 = notJS.GetElementByID("contactEditAddress1"); presenter.contactEditAddress1 == null {
		err = errors.New("unable to find #contactEditAddress1")
		return
	}
	if presenter.contactEditAddress2 = notJS.GetElementByID("contactEditAddress2"); presenter.contactEditAddress2 == null {
		err = errors.New("unable to find #contactEditAddress2")
		return
	}
	if presenter.contactEditCity = notJS.GetElementByID("contactEditCity"); presenter.contactEditCity == null {
		err = errors.New("unable to find #contactEditCity")
		return
	}
	if presenter.contactEditState = notJS.GetElementByID("contactEditState"); presenter.contactEditState == null {
		err = errors.New("unable to find #contactEditState")
		return
	}
	if presenter.contactEditZip = notJS.GetElementByID("contactEditZip"); presenter.contactEditZip == null {
		err = errors.New("unable to find #contactEditZip")
		return
	}
	if presenter.contactEditPhone = notJS.GetElementByID("contactEditPhone"); presenter.contactEditPhone == null {
		err = errors.New("unable to find #contactEditPhone")
		return
	}
	if presenter.contactEditEmail = notJS.GetElementByID("contactEditEmail"); presenter.contactEditEmail == null {
		err = errors.New("unable to find #contactEditEmail")
		return
	}
	if presenter.contactEditSocial = notJS.GetElementByID("contactEditSocial"); presenter.contactEditSocial == null {
		err = errors.New("unable to find #contactEditSocial")
		return
	}

	return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.

*/

func (presenter *panelPresenter) fillForm(r *record.Contact) {
	notJS.SetValue(presenter.contactEditName, r.Name)
	notJS.SetValue(presenter.contactEditAddress1, r.Address1)
	notJS.SetValue(presenter.contactEditAddress2, r.Address2)
	notJS.SetValue(presenter.contactEditCity, r.City)
	notJS.SetValue(presenter.contactEditState, r.State)
	notJS.SetValue(presenter.contactEditZip, r.Zip)
	notJS.SetValue(presenter.contactEditPhone, r.Phone)
	notJS.SetValue(presenter.contactEditEmail, r.Email)
	notJS.SetValue(presenter.contactEditSocial, r.Social)
}

func (presenter *panelPresenter) clearForm() {
	notJS.ClearValue(presenter.contactEditName)
	notJS.ClearValue(presenter.contactEditAddress1)
	notJS.ClearValue(presenter.contactEditAddress2)
	notJS.ClearValue(presenter.contactEditCity)
	notJS.ClearValue(presenter.contactEditState)
	notJS.ClearValue(presenter.contactEditZip)
	notJS.ClearValue(presenter.contactEditPhone)
	notJS.ClearValue(presenter.contactEditEmail)
	notJS.ClearValue(presenter.contactEditSocial)
}
