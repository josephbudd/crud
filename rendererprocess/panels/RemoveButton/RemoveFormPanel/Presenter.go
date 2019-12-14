// +build js, wasm

package removeformpanel

import (
	"github.com/josephbudd/crud/domain/store/record"
	"github.com/josephbudd/crud/rendererprocess/api/markup"
	"github.com/pkg/errors"
)

/*

	Panel name: RemoveFormPanel

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

	contactRemoveName     *markup.Element
	contactRemoveAddress1 *markup.Element
	contactRemoveAddress2 *markup.Element
	contactRemoveCity     *markup.Element
	contactRemoveState    *markup.Element
	contactRemoveZip      *markup.Element
	contactRemovePhone    *markup.Element
	contactRemoveEmail    *markup.Element
	contactRemoveSocial   *markup.Element
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

	// example:

	// Define the edit form's customer name input field.
	if presenter.editCustomerName = document.ElementByID("editCustomerName"); presenter.editCustomerName == nil {
		err = errors.New("unable to find #editCustomerName")
		return
	}

	*/

	if presenter.contactRemoveName = document.ElementByID("contactRemoveName"); presenter.contactRemoveName == nil {
		err = errors.New("unable to find #contactRemoveName")
		return
	}
	if presenter.contactRemoveAddress1 = document.ElementByID("contactRemoveAddress1"); presenter.contactRemoveAddress1 == nil {
		err = errors.New("unable to find #contactRemoveAddress1")
		return
	}
	if presenter.contactRemoveAddress2 = document.ElementByID("contactRemoveAddress2"); presenter.contactRemoveAddress2 == nil {
		err = errors.New("unable to find #contactRemoveAddress2")
		return
	}
	if presenter.contactRemoveCity = document.ElementByID("contactRemoveCity"); presenter.contactRemoveCity == nil {
		err = errors.New("unable to find #contactRemoveCity")
		return
	}
	if presenter.contactRemoveState = document.ElementByID("contactRemoveState"); presenter.contactRemoveState == nil {
		err = errors.New("unable to find #contactRemoveState")
		return
	}
	if presenter.contactRemoveZip = document.ElementByID("contactRemoveZip"); presenter.contactRemoveZip == nil {
		err = errors.New("unable to find #contactRemoveZip")
		return
	}
	if presenter.contactRemovePhone = document.ElementByID("contactRemovePhone"); presenter.contactRemovePhone == nil {
		err = errors.New("unable to find #contactRemovePhone")
		return
	}
	if presenter.contactRemoveEmail = document.ElementByID("contactRemoveEmail"); presenter.contactRemoveEmail == nil {
		err = errors.New("unable to find #contactRemoveEmail")
		return
	}
	if presenter.contactRemoveSocial = document.ElementByID("contactRemoveSocial"); presenter.contactRemoveSocial == nil {
		err = errors.New("unable to find #contactRemoveSocial")
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
	presenter.contactRemoveName.SetInnerText(r.Name)
	presenter.contactRemoveAddress1.SetInnerText(r.Address1)
	presenter.contactRemoveAddress2.SetInnerText(r.Address2)
	presenter.contactRemoveCity.SetInnerText(r.City)
	presenter.contactRemoveState.SetInnerText(r.State)
	presenter.contactRemoveZip.SetInnerText(r.Zip)
	presenter.contactRemovePhone.SetInnerText(r.Phone)
	presenter.contactRemoveEmail.SetInnerText(r.Email)
	presenter.contactRemoveSocial.SetInnerText(r.Social)
}

func (presenter *panelPresenter) clearForm() {
	presenter.contactRemoveName.SetInnerText(emptyText)
	presenter.contactRemoveAddress1.SetInnerText(emptyText)
	presenter.contactRemoveAddress2.SetInnerText(emptyText)
	presenter.contactRemoveCity.SetInnerText(emptyText)
	presenter.contactRemoveState.SetInnerText(emptyText)
	presenter.contactRemoveZip.SetInnerText(emptyText)
	presenter.contactRemovePhone.SetInnerText(emptyText)
	presenter.contactRemoveEmail.SetInnerText(emptyText)
	presenter.contactRemoveSocial.SetInnerText(emptyText)
}
