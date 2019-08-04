package removeselectpanel

import (
	"github.com/pkg/errors"

	"github.com/josephbudd/crud/renderer/paneling"
	"github.com/josephbudd/crud/renderer/lpc"
	"github.com/josephbudd/crud/renderer/notjs"
	"github.com/josephbudd/crud/renderer/viewtools"
)

/*

	Panel name: RemoveSelectPanel

*/

// Panel has a controler, presenter and caller.
// It also has show panel funcs for each panel in this panel group.
type Panel struct {
	controler *panelControler
	presenter *panelPresenter
	caller    *panelCaller
}

// NewPanel constructs a new panel.
func NewPanel(quitChan, eojChan chan struct{}, receiveChan lpc.Receiving, sendChan lpc.Sending, vtools *viewtools.Tools, njs *notjs.NotJS, help *paneling.Help) (panel *Panel, err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "RemoveSelectPanel")
		}
	}()

	quitCh = quitChan
	eojCh = eojChan
	receiveCh = receiveChan
	sendCh = sendChan
	tools = vtools
	notJS = njs

	group := &panelGroup{}
	controler := &panelControler{
		group: group,
	}
	presenter := &panelPresenter{
		group:          group,
	}
	caller := &panelCaller{
		group: group,
	}

	/* NOTE TO DEVELOPER. Step 1 of 1.

	// Set any controler, presenter or caller members that you added.
	// Use your custom help funcs if needed.
	// example:

	caller.state = help.GetStateAdd()
	
	*/

	controler.presenter = presenter
	controler.caller = caller
	presenter.controler = controler
	presenter.caller = caller
	caller.controler = controler
	caller.presenter = presenter

	// completions
	if err = group.defineMembers(); err != nil {
		return
	}
	if err = controler.defineControlsSetHandlers(); err != nil {
		return
	}
	if err = presenter.defineMembers(); err != nil {
		return
	}

	// No errors so define the panel.
	panel = &Panel{
		controler: controler,
		presenter: presenter,
		caller:     caller,
	}
	return
}

// Listen get the panel caller listening for main process messages.
func (panel *Panel) Listen() {
	panel.caller.listen()
}

// InitialCalls runs the first code that the panel needs to run.
func (panel *Panel) InitialCalls() {
	panel.controler.initialCalls()
	panel.caller.initialCalls()
}
