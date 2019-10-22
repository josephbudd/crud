// +build js, wasm

package addpanel

import (
	"syscall/js"

	"github.com/pkg/errors"
)

/*

	Panel name: AddPanel

*/

// panelPresenter writes to the panel
type panelPresenter struct {
	group      *panelGroup
	controller *panelController
	messenger  *panelMessenger

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your panelPresenter members here.

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

	if presenter.contactAddName = notJS.GetElementByID("contactAddName"); presenter.contactAddName == null {
		err = errors.New("unable to find #contactAddName")
		return
	}
	if presenter.contactAddAddress1 = notJS.GetElementByID("contactAddAddress1"); presenter.contactAddAddress1 == null {
		err = errors.New("unable to find #contactAddAddress1")
		return
	}
	if presenter.contactAddAddress2 = notJS.GetElementByID("contactAddAddress2"); presenter.contactAddAddress2 == null {
		err = errors.New("unable to find #contactAddAddress2")
		return
	}
	if presenter.contactAddCity = notJS.GetElementByID("contactAddCity"); presenter.contactAddCity == null {
		err = errors.New("unable to find #contactAddCity")
		return
	}
	if presenter.contactAddState = notJS.GetElementByID("contactAddState"); presenter.contactAddState == null {
		err = errors.New("unable to find #contactAddState")
		return
	}
	if presenter.contactAddZip = notJS.GetElementByID("contactAddZip"); presenter.contactAddZip == null {
		err = errors.New("unable to find #contactAddZip")
		return
	}
	if presenter.contactAddPhone = notJS.GetElementByID("contactAddPhone"); presenter.contactAddPhone == null {
		err = errors.New("unable to find #contactAddPhone")
		return
	}
	if presenter.contactAddEmail = notJS.GetElementByID("contactAddEmail"); presenter.contactAddEmail == null {
		err = errors.New("unable to find #contactAddEmail")
		return
	}
	if presenter.contactAddSocial = notJS.GetElementByID("contactAddSocial"); presenter.contactAddSocial == null {
		err = errors.New("unable to find #contactAddSocial")
		return
	}

	return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.

*/

func (presenter *panelPresenter) clearForm() {
	notJS.ClearValue(presenter.contactAddName)
	notJS.ClearValue(presenter.contactAddAddress1)
	notJS.ClearValue(presenter.contactAddAddress2)
	notJS.ClearValue(presenter.contactAddCity)
	notJS.ClearValue(presenter.contactAddState)
	notJS.ClearValue(presenter.contactAddZip)
	notJS.ClearValue(presenter.contactAddPhone)
	notJS.ClearValue(presenter.contactAddEmail)
	notJS.ClearValue(presenter.contactAddSocial)
}
