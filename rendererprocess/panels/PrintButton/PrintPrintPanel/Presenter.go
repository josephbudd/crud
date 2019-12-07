// +build js, wasm

package printprintpanel

import (
	"fmt"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/josephbudd/crud/rendererprocess/markup"
	"github.com/pkg/errors"
)

/*

	Panel name: PrintPrintPanel

*/

// panelPresenter writes to the panel
type panelPresenter struct {
	group      *panelGroup
	controller *panelController
	messenger  *panelMessenger

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your panelPresenter members here.

	// example:

	import "github.com/josephbudd/crud/rendererprocess/markup"

	editCustomerName *markup.Element

	*/

	printTitle   *markup.Element
	printTable   *markup.Element
	printPrintAs *markup.Element
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

	if presenter.printTitle = document.ElementByID("printTitle"); presenter.printTitle == nil {
		err = errors.New("unable to find #printTitle")
		return
	}
	if presenter.printTable = document.ElementByID("printTable"); presenter.printTable == nil {
		err = errors.New("unable to find #printTable")
		return
	}
	if presenter.printPrintAs = document.ElementByID("printPrintAs"); presenter.printPrintAs == nil {
		err = errors.New("unable to find #printPrintAs")
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

func (presenter *panelPresenter) displayContact(r *record.Contact) {
	// Store the print title for the controller's func handlePrint.
	documentPrintTitle = fmt.Sprintf("Contact: %s", r.Name)
	// Display the print heading.
	presenter.printTitle.SetInnerText(documentPrintTitle)
	// Set the print button label.
	presenter.printPrintAs.SetInnerText(fmt.Sprintf(`Print as "Contact: %s"`, r.Name))
	// The contact record will be displayed in the table.
	presenter.printTable.RemoveChildren()
	tbody := document.NewTBODY()
	presenter.newLine(tbody, "Name", r.Name)
	presenter.newLine(tbody, "Address 1", r.Address1)
	if len(r.Address2) > 0 {
		presenter.newLine(tbody, "Address 2", r.Address2)
	}
	presenter.newLine(tbody, "City", r.City)
	presenter.newLine(tbody, "State", r.State)
	presenter.newLine(tbody, "Zip", r.Zip)
	presenter.newLine(tbody, "Phone", r.Phone)
	presenter.newLine(tbody, "Email", r.Email)
	presenter.newLine(tbody, "Social", r.Social)
	presenter.printTable.AppendChild(tbody)
	// Show this panel.
	presenter.group.showPrintPrintPanel(false)
}

func (presenter *panelPresenter) newLine(tbody *markup.Element, fieldLabel, fieldText string) {
	tr := document.NewTR()
	th := document.NewTH()
	th.AddClass("label")
	th.SetInnerText(fieldLabel)
	tr.AppendChild(th)
	td := document.NewTD()
	td.AddClass("text")
	td.SetInnerText(fieldText)
	tr.AppendChild(td)
	tbody.AppendChild(tr)
}
