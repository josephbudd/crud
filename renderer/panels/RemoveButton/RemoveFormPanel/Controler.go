package removeformpanel

import (
	"syscall/js"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/pkg/errors"
)

/*

	Panel name: RemoveFormPanel

*/

// panelControler controls user input.
type panelControler struct {
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelControler members.

	*/

	contact             *record.Contact
	contactRemoveSubmit js.Value
	contactRemoveCancel js.Value
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

	if controler.contactRemoveSubmit = notJS.GetElementByID("contactRemoveSubmit"); controler.contactRemoveSubmit == null {
		err = errors.New("unable to find #contactRemoveSubmit")
		return
	}
	cb := tools.RegisterEventCallBack(controler.handleSubmit, true, true, true)
	notJS.SetOnClick(controler.contactRemoveSubmit, cb)

	if controler.contactRemoveCancel = notJS.GetElementByID("contactRemoveCancel"); controler.contactRemoveCancel == null {
		err = errors.New("unable to find #contactRemoveCancel")
		return
	}
	cb = tools.RegisterEventCallBack(controler.handleCancel, true, true, true)
	notJS.SetOnClick(controler.contactRemoveCancel, cb)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (controler *panelControler) handleSubmit(event js.Value) (nilReturn interface{}) {
	controler.caller.removeContact(controler.contact.ID)
	return
}

func (controler *panelControler) handleCancel(event js.Value) (nilReturn interface{}) {
	controler.presenter.clearForm()
	// Go back to the RemoveSelectPanel
	controler.group.showRemoveSelectPanel(false)
	return
}

func (controler *panelControler) handleGetContact(r *record.Contact) {
	controler.contact = r
	controler.presenter.fillForm(r)
	controler.group.showRemoveFormPanel(false)
}

// initialCalls runs the first code that the controler needs to run.
func (controler *panelControler) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/

}
