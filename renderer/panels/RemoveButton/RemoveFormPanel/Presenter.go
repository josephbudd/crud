package removeformpanel

import (
	"syscall/js"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/pkg/errors"
)

/*

	Panel name: RemoveFormPanel

*/

// panelPresenter writes to the panel
type panelPresenter struct {
	group     *panelGroup
	controler *panelControler
	caller    *panelCaller

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your panelPresenter members here.

	*/

	contactRemoveName     js.Value
	contactRemoveAddress1 js.Value
	contactRemoveAddress2 js.Value
	contactRemoveCity     js.Value
	contactRemoveState    js.Value
	contactRemoveZip      js.Value
	contactRemovePhone    js.Value
	contactRemoveEmail    js.Value
	contactRemoveSocial   js.Value
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

	if presenter.contactRemoveName = notJS.GetElementByID("contactRemoveName"); presenter.contactRemoveName == null {
		err = errors.New("unable to find #contactRemoveName")
		return
	}
	if presenter.contactRemoveAddress1 = notJS.GetElementByID("contactRemoveAddress1"); presenter.contactRemoveAddress1 == null {
		err = errors.New("unable to find #contactRemoveAddress1")
		return
	}
	if presenter.contactRemoveAddress2 = notJS.GetElementByID("contactRemoveAddress2"); presenter.contactRemoveAddress2 == null {
		err = errors.New("unable to find #contactRemoveAddress2")
		return
	}
	if presenter.contactRemoveCity = notJS.GetElementByID("contactRemoveCity"); presenter.contactRemoveCity == null {
		err = errors.New("unable to find #contactRemoveCity")
		return
	}
	if presenter.contactRemoveState = notJS.GetElementByID("contactRemoveState"); presenter.contactRemoveState == null {
		err = errors.New("unable to find #contactRemoveState")
		return
	}
	if presenter.contactRemoveZip = notJS.GetElementByID("contactRemoveZip"); presenter.contactRemoveZip == null {
		err = errors.New("unable to find #contactRemoveZip")
		return
	}
	if presenter.contactRemovePhone = notJS.GetElementByID("contactRemovePhone"); presenter.contactRemovePhone == null {
		err = errors.New("unable to find #contactRemovePhone")
		return
	}
	if presenter.contactRemoveEmail = notJS.GetElementByID("contactRemoveEmail"); presenter.contactRemoveEmail == null {
		err = errors.New("unable to find #contactRemoveEmail")
		return
	}
	if presenter.contactRemoveSocial = notJS.GetElementByID("contactRemoveSocial"); presenter.contactRemoveSocial == null {
		err = errors.New("unable to find #contactRemoveSocial")
		return
	}

	return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.

*/

func (presenter *panelPresenter) fillForm(r *record.Contact) {
	notJS.SetInnerText(presenter.contactRemoveName, r.Name)
	notJS.SetInnerText(presenter.contactRemoveAddress1, r.Address1)
	notJS.SetInnerText(presenter.contactRemoveAddress2, r.Address2)
	notJS.SetInnerText(presenter.contactRemoveCity, r.City)
	notJS.SetInnerText(presenter.contactRemoveState, r.State)
	notJS.SetInnerText(presenter.contactRemoveZip, r.Zip)
	notJS.SetInnerText(presenter.contactRemovePhone, r.Phone)
	notJS.SetInnerText(presenter.contactRemoveEmail, r.Email)
	notJS.SetInnerText(presenter.contactRemoveSocial, r.Social)
}

func (presenter *panelPresenter) clearForm() {
	notJS.SetInnerText(presenter.contactRemoveName, emptyString)
	notJS.SetInnerText(presenter.contactRemoveAddress1, emptyString)
	notJS.SetInnerText(presenter.contactRemoveAddress2, emptyString)
	notJS.SetInnerText(presenter.contactRemoveCity, emptyString)
	notJS.SetInnerText(presenter.contactRemoveState, emptyString)
	notJS.SetInnerText(presenter.contactRemoveZip, emptyString)
	notJS.SetInnerText(presenter.contactRemovePhone, emptyString)
	notJS.SetInnerText(presenter.contactRemoveEmail, emptyString)
	notJS.SetInnerText(presenter.contactRemoveSocial, emptyString)
}
