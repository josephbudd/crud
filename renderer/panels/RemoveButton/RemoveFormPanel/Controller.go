package removeformpanel

import (
	"syscall/js"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/pkg/errors"
)

/*

	Panel name: RemoveFormPanel

*/

// panelController controls user input.
type panelController struct {
	group     *panelGroup
	presenter *panelPresenter
	caller    *panelCaller

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelController members.

	*/

	contact             *record.Contact
	contactRemoveSubmit js.Value
	contactRemoveCancel js.Value
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

	if controller.contactRemoveSubmit = notJS.GetElementByID("contactRemoveSubmit"); controller.contactRemoveSubmit == null {
		err = errors.New("unable to find #contactRemoveSubmit")
		return
	}
	cb := tools.RegisterEventCallBack(controller.handleSubmit, true, true, true)
	notJS.SetOnClick(controller.contactRemoveSubmit, cb)

	if controller.contactRemoveCancel = notJS.GetElementByID("contactRemoveCancel"); controller.contactRemoveCancel == null {
		err = errors.New("unable to find #contactRemoveCancel")
		return
	}
	cb = tools.RegisterEventCallBack(controller.handleCancel, true, true, true)
	notJS.SetOnClick(controller.contactRemoveCancel, cb)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (controller *panelController) handleSubmit(event js.Value) (nilReturn interface{}) {
	controller.caller.removeContact(controller.contact.ID)
	return
}

func (controller *panelController) handleCancel(event js.Value) (nilReturn interface{}) {
	controller.presenter.clearForm()
	// Go back to the RemoveSelectPanel
	controller.group.showRemoveSelectPanel(false)
	return
}

func (controller *panelController) handleGetContact(r *record.Contact) {
	controller.contact = r
	controller.presenter.fillForm(r)
	controller.group.showRemoveFormPanel(false)
}

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	*/

}
