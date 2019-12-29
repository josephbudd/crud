// +build js, wasm

package editformpanel

import (
	"fmt"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/josephbudd/crud/rendererprocess/api/markup"
)

/*

	Panel name: EditFormPanel

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

	contactEditName     *markup.Element
	contactEditAddress1 *markup.Element
	contactEditAddress2 *markup.Element
	contactEditCity     *markup.Element
	contactEditState    *markup.Element
	contactEditZip      *markup.Element
	contactEditPhone    *markup.Element
	contactEditEmail    *markup.Element
	contactEditSocial   *markup.Element
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

	if presenter.contactEditName = document.ElementByID("contactEditName"); presenter.contactEditName == nil {
		err = fmt.Errorf("unable to find #contactEditName")
		return
	}
	if presenter.contactEditAddress1 = document.ElementByID("contactEditAddress1"); presenter.contactEditAddress1 == nil {
		err = fmt.Errorf("unable to find #contactEditAddress1")
		return
	}
	if presenter.contactEditAddress2 = document.ElementByID("contactEditAddress2"); presenter.contactEditAddress2 == nil {
		err = fmt.Errorf("unable to find #contactEditAddress2")
		return
	}
	if presenter.contactEditCity = document.ElementByID("contactEditCity"); presenter.contactEditCity == nil {
		err = fmt.Errorf("unable to find #contactEditCity")
		return
	}
	if presenter.contactEditState = document.ElementByID("contactEditState"); presenter.contactEditState == nil {
		err = fmt.Errorf("unable to find #contactEditState")
		return
	}
	if presenter.contactEditZip = document.ElementByID("contactEditZip"); presenter.contactEditZip == nil {
		err = fmt.Errorf("unable to find #contactEditZip")
		return
	}
	if presenter.contactEditPhone = document.ElementByID("contactEditPhone"); presenter.contactEditPhone == nil {
		err = fmt.Errorf("unable to find #contactEditPhone")
		return
	}
	if presenter.contactEditEmail = document.ElementByID("contactEditEmail"); presenter.contactEditEmail == nil {
		err = fmt.Errorf("unable to find #contactEditEmail")
		return
	}
	if presenter.contactEditSocial = document.ElementByID("contactEditSocial"); presenter.contactEditSocial == nil {
		err = fmt.Errorf("unable to find #contactEditSocial")
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

func (presenter *panelPresenter) fillForm(r *record.Contact) {
	presenter.contactEditName.SetValue(r.Name)
	presenter.contactEditAddress1.SetValue(r.Address1)
	presenter.contactEditAddress2.SetValue(r.Address2)
	presenter.contactEditCity.SetValue(r.City)
	presenter.contactEditState.SetValue(r.State)
	presenter.contactEditZip.SetValue(r.Zip)
	presenter.contactEditPhone.SetValue(r.Phone)
	presenter.contactEditEmail.SetValue(r.Email)
	presenter.contactEditSocial.SetValue(r.Social)
}

func (presenter *panelPresenter) clearForm() {
	presenter.contactEditName.ClearValue()
	presenter.contactEditAddress1.ClearValue()
	presenter.contactEditAddress2.ClearValue()
	presenter.contactEditCity.ClearValue()
	presenter.contactEditState.ClearValue()
	presenter.contactEditZip.ClearValue()
	presenter.contactEditPhone.ClearValue()
	presenter.contactEditEmail.ClearValue()
	presenter.contactEditSocial.ClearValue()
}
