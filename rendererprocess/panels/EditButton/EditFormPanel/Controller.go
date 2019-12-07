// +build js, wasm

package editformpanel

import (
	"github.com/josephbudd/crud/domain/store/record"
	"github.com/josephbudd/crud/rendererprocess/display"
	"github.com/josephbudd/crud/rendererprocess/event"
	"github.com/josephbudd/crud/rendererprocess/markup"
	"github.com/pkg/errors"
)

/*

	Panel name: EditFormPanel

*/

// panelController controls user input.
type panelController struct {
	group     *panelGroup
	presenter *panelPresenter
	messenger *panelMessenger

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// Declare your panelController fields.

	// example:

	import "syscall/js"
	import "github.com/josephbudd/crud/rendererprocess/markup"

	addCustomerName   *markup.Element
	addCustomerSubmit *markup.Element

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
	contactEditSubmit   *markup.Element
	contactEditCancel   *markup.Element

	contact *record.Contact
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

	// example:

	// Define the customer name text input GUI controller.
	if controller.addCustomerName = document.ElementByID("addCustomerName"); controller.addCustomerName == nil {
		err = errors.New("unable to find #addCustomerName")
		return
	}

	// Define the submit button GUI controller.
	if controller.addCustomerSubmit = document.ElementByID("addCustomerSubmit"); controller.addCustomerSubmit == nil {
		err = errors.New("unable to find #addCustomerSubmit")
		return
	}
	// Handle the submit button's onclick event.
	controller.addCustomerSubmit.SetEventHandler(controller.handleSubmit, "click", false)

	*/

	if controller.contactEditName = document.ElementByID("contactEditName"); controller.contactEditName == nil {
		err = errors.New("unable to find #contactEditName")
		return
	}
	if controller.contactEditAddress1 = document.ElementByID("contactEditAddress1"); controller.contactEditAddress1 == nil {
		err = errors.New("unable to find #contactEditAddress1")
		return
	}
	if controller.contactEditAddress2 = document.ElementByID("contactEditAddress2"); controller.contactEditAddress2 == nil {
		err = errors.New("unable to find #contactEditAddress2")
		return
	}
	if controller.contactEditCity = document.ElementByID("contactEditCity"); controller.contactEditCity == nil {
		err = errors.New("unable to find #contactEditCity")
		return
	}
	if controller.contactEditState = document.ElementByID("contactEditState"); controller.contactEditState == nil {
		err = errors.New("unable to find #contactEditState")
		return
	}
	if controller.contactEditZip = document.ElementByID("contactEditZip"); controller.contactEditZip == nil {
		err = errors.New("unable to find #contactEditZip")
		return
	}
	if controller.contactEditPhone = document.ElementByID("contactEditPhone"); controller.contactEditPhone == nil {
		err = errors.New("unable to find #contactEditPhone")
		return
	}
	if controller.contactEditEmail = document.ElementByID("contactEditEmail"); controller.contactEditEmail == nil {
		err = errors.New("unable to find #contactEditEmail")
		return
	}
	if controller.contactEditSocial = document.ElementByID("contactEditSocial"); controller.contactEditSocial == nil {
		err = errors.New("unable to find #contactEditSocial")
		return
	}

	if controller.contactEditSubmit = document.ElementByID("contactEditSubmit"); controller.contactEditSubmit == nil {
		err = errors.New("unable to find #contactEditSubmit")
		return
	}
	// Handle the edit button's onclick event.
	controller.contactEditSubmit.SetEventHandler(controller.handleSubmit, "click", false)

	if controller.contactEditCancel = document.ElementByID("contactEditCancel"); controller.contactEditCancel == nil {
		err = errors.New("unable to find #contactEditCancel")
		return
	}
	// Handle the cancel button's onclick event.
	controller.contactEditCancel.SetEventHandler(controller.handleCancel, "click", false)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

// example:

import "github.com/josephbudd/crud/domain/store/record"
import "github.com/josephbudd/crud/rendererprocess/event"
import "github.com/josephbudd/crud/rendererprocess/display"

func (controller *panelController) handleSubmit(e event.Event) (nilReturn interface{}) {
	// See renderer/event/event.go.
	// The event.Event funcs.
	//   e.PreventDefaultBehavior()
	//   e.StopCurrentPhasePropagation()
	//   e.StopAllPhasePropagation()
	//   target := e.JSTarget
	//   event := e.JSEvent
	// You must use the javascript event e.JSEvent, as a js.Value.
	// However, you can use the target as a *markup.Element
	//   target := document.NewElementFromJSValue(e.JSTarget)

	name := strings.TrimSpace(controller.addCustomerName.Value())
	if len(name) == 0 {
		display.Error("Customer Name is required.")
		return
	}
	r := &record.Customer{
		Name: name,
	}
	controller.messenger.AddCustomer(r)
	return
}

*/

func (controller *panelController) handleSubmit(e event.Event) (nilReturn interface{}) {
	r := controller.getRecord()
	if len(r.Name) == 0 {
		display.Error("Name is required.")
		return
	}
	if len(r.Address1) == 0 {
		display.Error("Address1 is required.")
		return
	}
	if len(r.City) == 0 {
		display.Error("City is required.")
		return
	}
	if len(r.State) == 0 {
		display.Error("State is required.")
		return
	}
	if len(r.Zip) == 0 {
		display.Error("Zip is required.")
		return
	}
	if len(r.Email) == 0 && len(r.Phone) == 0 {
		display.Error("Either Email or Phone is required.")
		return
	}
	controller.messenger.editContact(r)
	return
}

func (controller *panelController) handleCancel(e event.Event) (nilReturn interface{}) {
	controller.presenter.clearForm()
	// Go back to the EditSelectPanel
	controller.group.showEditSelectPanel(false)
	return
}

func (controller *panelController) handleGetContact(r *record.Contact) {
	controller.contact = r
	controller.presenter.fillForm(r)
	controller.group.showEditFormPanel(false)
}

func (controller *panelController) getRecord() (r *record.Contact) {
	r = &record.Contact{
		ID:       controller.contact.ID,
		Name:     controller.contactEditName.Value(),
		Address1: controller.contactEditAddress1.Value(),
		Address2: controller.contactEditAddress2.Value(),
		City:     controller.contactEditCity.Value(),
		State:    controller.contactEditState.Value(),
		Zip:      controller.contactEditZip.Value(),
		Phone:    controller.contactEditPhone.Value(),
		Email:    controller.contactEditEmail.Value(),
		Social:   controller.contactEditSocial.Value(),
	}
	return
}

// initialCalls runs the first code that the controller needs to run.
func (controller *panelController) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	// Make the initial calls.
	// I use this to start up widgets. For example a virtual list widget.

	// example:

	controller.customerSelectWidget.start()

	*/
}
