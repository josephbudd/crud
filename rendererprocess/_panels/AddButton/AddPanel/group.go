// +build js, wasm

package addpanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/crud/rendererprocess/markup"
	"github.com/josephbudd/crud/rendererprocess/framework/viewtools"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

// panelGroup is a group of 1 panel.
// It also has a show panel func for each panel in this panel group.
type panelGroup struct {
	addPanel js.Value
}

func (group *panelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(group *panelGroup) defineMembers()")
		}
	}()

    var panel *markup.Element
 if panel = document.ElementByID("mainMasterView-home-pad-AddButton-AddPanel"); panel == nil {
		err = errors.New("unable to find #mainMasterView-home-pad-AddButton-AddPanel")
		return
    }
    group.addPanel = panel.JSValue()

	return
}

/*
	Show panel funcs.

	Call these from the controller, presenter and messenger.
*/

// showAddPanel shows the panel you named AddPanel while hiding any other panels in this panel group.
// This panel's id is mainMasterView-home-pad-AddButton-AddPanel.
// This panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when this panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for this panel is:
Presentation:
  An empty contact form, "Add" and "Cancel" buttons.
  * An add button that either warns of missing fields or submits the form.
  * A cancel button that clears the form and goes back.
  * A success or error alert to the user when the main process reports back on the submittion.
Render:
  Add button handler warns of missing fields or submits the form input as a new contact record.
Main Process:
  Saves the new record updating the record id. Reports back to the renderer with updated record and error.
LPC:
  AddContact
  RendererToMainProcess:
    Record: Contact Record
  MainProcessToRenderer;
    Error: bool
    ErrorMessage: string
    Record: Contact Record

*/
func (group *panelGroup) showAddPanel(force bool) {
	viewtools.ShowPanelInButtonGroup(group.addPanel, force)
}
