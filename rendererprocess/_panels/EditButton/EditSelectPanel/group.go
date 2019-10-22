// +build js, wasm

package editselectpanel

import (
	"syscall/js"

	"github.com/pkg/errors"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

// panelGroup is a group of 3 panels.
// It also has show panel funcs for each panel in this panel group.
type panelGroup struct {
	editFormPanel js.Value
	editNotReadyPanel js.Value
	editSelectPanel js.Value
}

func (group *panelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(group *panelGroup) defineMembers()")
		}
	}()

	if group.editFormPanel = notJS.GetElementByID("tabsMasterView-home-pad-EditButton-EditFormPanel"); group.editFormPanel == null {
		err = errors.New("unable to find #tabsMasterView-home-pad-EditButton-EditFormPanel")
		return
	}
	if group.editNotReadyPanel = notJS.GetElementByID("tabsMasterView-home-pad-EditButton-EditNotReadyPanel"); group.editNotReadyPanel == null {
		err = errors.New("unable to find #tabsMasterView-home-pad-EditButton-EditNotReadyPanel")
		return
	}
	if group.editSelectPanel = notJS.GetElementByID("tabsMasterView-home-pad-EditButton-EditSelectPanel"); group.editSelectPanel == null {
		err = errors.New("unable to find #tabsMasterView-home-pad-EditButton-EditSelectPanel")
		return
	}

	return
}

/*
	Show panel funcs.

	Call these from the controller, presenter and messenger.
*/

// showEditFormPanel shows the panel you named EditFormPanel while hiding any other panels in this panel group.
// That panel's id is tabsMasterView-home-pad-EditButton-EditFormPanel.
// That panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when that panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for that panel is:
A contact form with "edit" and "cancel" buttons.
*/
func (group *panelGroup) showEditFormPanel(force bool) {
	tools.ShowPanelInButtonGroup(group.editFormPanel, force)
}

// showEditNotReadyPanel shows the panel you named EditNotReadyPanel while hiding any other panels in this panel group.
// That panel's id is tabsMasterView-home-pad-EditButton-EditNotReadyPanel.
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
	tools.ShowPanelInButtonGroup(group.editNotReadyPanel, force)
}

// showEditSelectPanel shows the panel you named EditSelectPanel while hiding any other panels in this panel group.
// This panel's id is tabsMasterView-home-pad-EditButton-EditSelectPanel.
// This panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when this panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for this panel is:
A select widget sorted by name.
The widget needs to be rebuilt anytime a conact is added, edited or removed.

*/
func (group *panelGroup) showEditSelectPanel(force bool) {
	tools.ShowPanelInButtonGroup(group.editSelectPanel, force)
}
