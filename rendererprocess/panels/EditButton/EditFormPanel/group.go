// +build js, wasm

package editformpanel

import (
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/crud/rendererprocess/api/markup"
	"github.com/josephbudd/crud/rendererprocess/framework/viewtools"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

// panelGroup is a group of 3 panels.
// It also has show panel funcs for each panel in this panel group.
type panelGroup struct {
	editSelectPanel js.Value
	editFormPanel js.Value
	editNotReadyPanel js.Value
}

func (group *panelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(group *panelGroup) defineMembers()")
		}
	}()

    var panel *markup.Element
 if panel = document.ElementByID("mainMasterView-home-pad-EditButton-EditSelectPanel"); panel == nil {
		err = errors.New("unable to find #mainMasterView-home-pad-EditButton-EditSelectPanel")
		return
    }
    group.editSelectPanel = panel.JSValue()
 if panel = document.ElementByID("mainMasterView-home-pad-EditButton-EditFormPanel"); panel == nil {
		err = errors.New("unable to find #mainMasterView-home-pad-EditButton-EditFormPanel")
		return
    }
    group.editFormPanel = panel.JSValue()
 if panel = document.ElementByID("mainMasterView-home-pad-EditButton-EditNotReadyPanel"); panel == nil {
		err = errors.New("unable to find #mainMasterView-home-pad-EditButton-EditNotReadyPanel")
		return
    }
    group.editNotReadyPanel = panel.JSValue()

	return
}

/*
	Show panel funcs.

	Call these from the controller, presenter and messenger.
*/

// showEditSelectPanel shows the panel you named EditSelectPanel while hiding any other panels in this panel group.
// That panel's id is mainMasterView-home-pad-EditButton-EditSelectPanel.
// That panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when that panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for that panel is:
A select widget sorted by name.
The widget needs to be rebuilt anytime a conact is added, edited or removed.

*/
func (group *panelGroup) showEditSelectPanel(force bool) {
	viewtools.ShowPanelInButtonGroup(group.editSelectPanel, force)
}

// showEditFormPanel shows the panel you named EditFormPanel while hiding any other panels in this panel group.
// This panel's id is mainMasterView-home-pad-EditButton-EditFormPanel.
// This panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when this panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for this panel is:
A contact form with "edit" and "cancel" buttons.
*/
func (group *panelGroup) showEditFormPanel(force bool) {
	viewtools.ShowPanelInButtonGroup(group.editFormPanel, force)
}

// showEditNotReadyPanel shows the panel you named EditNotReadyPanel while hiding any other panels in this panel group.
// That panel's id is mainMasterView-home-pad-EditButton-EditNotReadyPanel.
// That panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when that panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for that panel is:
Static text displayed when there are no contacts in the store.
*/
func (group *panelGroup) showEditNotReadyPanel(force bool) {
	viewtools.ShowPanelInButtonGroup(group.editNotReadyPanel, force)
}
