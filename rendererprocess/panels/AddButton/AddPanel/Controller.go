// +build js, wasm

package addpanel

import (
	"github.com/josephbudd/crud/domain/store/record"
	"github.com/josephbudd/crud/rendererprocess/api/display"
	"github.com/josephbudd/crud/rendererprocess/api/event"
	"github.com/josephbudd/crud/rendererprocess/api/markup"
	"github.com/pkg/errors"
)

/*

	Panel name: AddPanel

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
	import "github.com/josephbudd/crud/rendererprocess/api/markup"

	addCustomerName   *markup.Element
	addCustomerSubmit *markup.Element

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
	contactAddSubmit   *markup.Element
	contactAddCancel   *markup.Element
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

	if controller.contactAddName = document.ElementByID("contactAddName"); controller.contactAddName == nil {
		err = errors.New("unable to find #contactAddName")
		return
	}
	if controller.contactAddAddress1 = document.ElementByID("contactAddAddress1"); controller.contactAddAddress1 == nil {
		err = errors.New("unable to find #contactAddAddress1")
		return
	}
	if controller.contactAddAddress2 = document.ElementByID("contactAddAddress2"); controller.contactAddAddress2 == nil {
		err = errors.New("unable to find #contactAddAddress2")
		return
	}
	if controller.contactAddCity = document.ElementByID("contactAddCity"); controller.contactAddCity == nil {
		err = errors.New("unable to find #contactAddCity")
		return
	}
	if controller.contactAddState = document.ElementByID("contactAddState"); controller.contactAddState == nil {
		err = errors.New("unable to find #contactAddState")
		return
	}
	if controller.contactAddZip = document.ElementByID("contactAddZip"); controller.contactAddZip == nil {
		err = errors.New("unable to find #contactAddZip")
		return
	}
	if controller.contactAddPhone = document.ElementByID("contactAddPhone"); controller.contactAddPhone == nil {
		err = errors.New("unable to find #contactAddPhone")
		return
	}
	if controller.contactAddEmail = document.ElementByID("contactAddEmail"); controller.contactAddEmail == nil {
		err = errors.New("unable to find #contactAddEmail")
		return
	}
	if controller.contactAddSocial = document.ElementByID("contactAddSocial"); controller.contactAddSocial == nil {
		err = errors.New("unable to find #contactAddSocial")
		return
	}

	if controller.contactAddSubmit = document.ElementByID("contactAddSubmit"); controller.contactAddSubmit == nil {
		err = errors.New("unable to find #contactAddSubmit")
		return
	}
	// Handle the add button's onclick event.
	controller.contactAddSubmit.SetEventHandler(controller.handleSubmit, "click", false)

	if controller.contactAddCancel = document.ElementByID("contactAddCancel"); controller.contactAddCancel == nil {
		err = errors.New("unable to find #contactAddCancel")
		return
	}
	// Handle the cancel button's onclick event.
	controller.contactAddCancel.SetEventHandler(controller.handleCancel, "click", false)

	return
}

/* NOTE TO DEVELOPER. Step 3 of 4.

// Handlers and other functions.

// example:

import "github.com/josephbudd/crud/domain/store/record"
import "github.com/josephbudd/crud/rendererprocess/api/event"
import "github.com/josephbudd/crud/rendererprocess/api/display"

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
	controller.messenger.addContact(r)
	return
}

func (controller *panelController) handleCancel(e event.Event) (nilReturn interface{}) {
	controller.presenter.clearForm()
	display.Back()
	return
}

func (controller *panelController) getRecord() (r *record.Contact) {
	r = &record.Contact{
		Name:     controller.contactAddName.Value(),
		Address1: controller.contactAddAddress1.Value(),
		Address2: controller.contactAddAddress2.Value(),
		City:     controller.contactAddCity.Value(),
		State:    controller.contactAddState.Value(),
		Zip:      controller.contactAddZip.Value(),
		Phone:    controller.contactAddPhone.Value(),
		Email:    controller.contactAddEmail.Value(),
		Social:   controller.contactAddSocial.Value(),
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
