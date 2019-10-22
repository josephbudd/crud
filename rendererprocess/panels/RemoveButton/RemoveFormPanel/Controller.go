// +build js, wasm

package removeformpanel

import (
	"log"
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/crud/domain/store/record"
	"github.com/josephbudd/crud/rendererprocess/viewtools"
)

/*

	Panel name: RemoveFormPanel

*/

// panelController controls user input.
type panelController struct {
	group     *panelGroup
	presenter *panelPresenter
	messenger *panelMessenger
	eventCh   chan viewtools.Event

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelController fields.

	*/

	contact             *record.Contact
	contactRemoveSubmit js.Value
	contactRemoveCancel js.Value
}

// defineControlsHandlers defines the GUI's controllers and their event handlers.
// Returns the error.
func (controller *panelController) defineControlsHandlers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(controller *panelController) defineControlsHandlers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 4.

	// Define each controller in the GUI by it's html element.
	// Handle each controller's events.

	*/

	if controller.contactRemoveSubmit = notJS.GetElementByID("contactRemoveSubmit"); controller.contactRemoveSubmit == null {
		err = errors.New("unable to find #contactRemoveSubmit")
		return
	}
	// Handle the submit button's onclick event.
	tools.AddEventHandler(controller.handleSubmit, controller.contactRemoveSubmit, "click", false)

	if controller.contactRemoveCancel = notJS.GetElementByID("contactRemoveCancel"); controller.contactRemoveCancel == null {
		err = errors.New("unable to find #contactRemoveCancel")
		return
	}
	// Handle the cancel button's onclick event.
	tools.AddEventHandler(controller.handleCancel, controller.contactRemoveCancel, "click", false)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

*/

func (controller *panelController) handleSubmit(e viewtools.Event) (nilReturn interface{}) {
	controller.messenger.removeContact(controller.contact.ID)
	return
}

func (controller *panelController) handleCancel(e viewtools.Event) (nilReturn interface{}) {
	controller.presenter.clearForm()
	// Go back to the RemoveSelectPanel
	controller.group.showRemoveSelectPanel(false)
	return
}

func (controller *panelController) handleGetContact(r *record.Contact) {
	log.Println("handleGetContact")
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
