// +build js, wasm

package addpanel

import (
	"fmt"

	"github.com/josephbudd/crud/rendererprocess/api/markup"
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

	// example:

	import "github.com/josephbudd/crud/rendererprocess/api/markup"

	editCustomerName *markup.Element

	*/

	contactAddName     *markup.Element
	contactAddAddress1 *markup.Element
	contactAddAddress2 *markup.Element
	contactAddCity     *markup.Element
	contactAddState    *markup.Element
	contactAddZip      *markup.Element
	contactAddPhone    *markup.Element
	contactAddEmail    *markup.Element
	contactAddSocial   *markup.Element
}

// defineMembers defines the panelPresenter members by their html elements.
// Returns the error.
func (presenter *panelPresenter) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("(presenter *panelPresenter) defineMembers(): %w", err)
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your panelPresenter members.

	// example:

	// Define the edit form's customer name input field.
	if presenter.editCustomerName = document.ElementByID("editCustomerName"); presenter.editCustomerName == nil {
		err = fmt.Errorf("unable to find #editCustomerName")
		return
	}

	*/

	if presenter.contactAddName = document.ElementByID("contactAddName"); presenter.contactAddName == nil {
		err = fmt.Errorf("unable to find #contactAddName")
		return
	}
	if presenter.contactAddAddress1 = document.ElementByID("contactAddAddress1"); presenter.contactAddAddress1 == nil {
		err = fmt.Errorf("unable to find #contactAddAddress1")
		return
	}
	if presenter.contactAddAddress2 = document.ElementByID("contactAddAddress2"); presenter.contactAddAddress2 == nil {
		err = fmt.Errorf("unable to find #contactAddAddress2")
		return
	}
	if presenter.contactAddCity = document.ElementByID("contactAddCity"); presenter.contactAddCity == nil {
		err = fmt.Errorf("unable to find #contactAddCity")
		return
	}
	if presenter.contactAddState = document.ElementByID("contactAddState"); presenter.contactAddState == nil {
		err = fmt.Errorf("unable to find #contactAddState")
		return
	}
	if presenter.contactAddZip = document.ElementByID("contactAddZip"); presenter.contactAddZip == nil {
		err = fmt.Errorf("unable to find #contactAddZip")
		return
	}
	if presenter.contactAddPhone = document.ElementByID("contactAddPhone"); presenter.contactAddPhone == nil {
		err = fmt.Errorf("unable to find #contactAddPhone")
		return
	}
	if presenter.contactAddEmail = document.ElementByID("contactAddEmail"); presenter.contactAddEmail == nil {
		err = fmt.Errorf("unable to find #contactAddEmail")
		return
	}
	if presenter.contactAddSocial = document.ElementByID("contactAddSocial"); presenter.contactAddSocial == nil {
		err = fmt.Errorf("unable to find #contactAddSocial")
		return
	}

	return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.

// example:

// displayCustomer displays the customer in the edit customer form panel.
func (presenter *panelPresenter) displayCustomer(record *types.CustomerRecord) {
	presenter.editCustomerName.SetValue(record.Name)
}

*/

func (presenter *panelPresenter) clearForm() {
	presenter.contactAddName.ClearValue()
	presenter.contactAddAddress1.ClearValue()
	presenter.contactAddAddress2.ClearValue()
	presenter.contactAddCity.ClearValue()
	presenter.contactAddState.ClearValue()
	presenter.contactAddZip.ClearValue()
	presenter.contactAddPhone.ClearValue()
	presenter.contactAddEmail.ClearValue()
	presenter.contactAddSocial.ClearValue()
}
