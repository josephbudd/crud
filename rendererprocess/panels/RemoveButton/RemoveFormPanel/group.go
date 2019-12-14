// +build js, wasm

package removeformpanel

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
	removeNotReadyPanel js.Value
	removeSelectPanel js.Value
	removeFormPanel js.Value
}

func (group *panelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(group *panelGroup) defineMembers()")
		}
	}()

    var panel *markup.Element
 if panel = document.ElementByID("mainMasterView-home-pad-RemoveButton-RemoveNotReadyPanel"); panel == nil {
		err = errors.New("unable to find #mainMasterView-home-pad-RemoveButton-RemoveNotReadyPanel")
		return
    }
    group.removeNotReadyPanel = panel.JSValue()
 if panel = document.ElementByID("mainMasterView-home-pad-RemoveButton-RemoveSelectPanel"); panel == nil {
		err = errors.New("unable to find #mainMasterView-home-pad-RemoveButton-RemoveSelectPanel")
		return
    }
    group.removeSelectPanel = panel.JSValue()
 if panel = document.ElementByID("mainMasterView-home-pad-RemoveButton-RemoveFormPanel"); panel == nil {
		err = errors.New("unable to find #mainMasterView-home-pad-RemoveButton-RemoveFormPanel")
		return
    }
    group.removeFormPanel = panel.JSValue()

	return
}

/*
	Show panel funcs.

	Call these from the controller, presenter and messenger.
*/

// showRemoveNotReadyPanel shows the panel you named RemoveNotReadyPanel while hiding any other panels in this panel group.
// That panel's id is mainMasterView-home-pad-RemoveButton-RemoveNotReadyPanel.
// That panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when that panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for that panel is:
Static text displayed when there are no contacts in the store.
*/
func (group *panelGroup) showRemoveNotReadyPanel(force bool) {
	viewtools.ShowPanelInButtonGroup(group.removeNotReadyPanel, force)
}

// showRemoveSelectPanel shows the panel you named RemoveSelectPanel while hiding any other panels in this panel group.
// That panel's id is mainMasterView-home-pad-RemoveButton-RemoveSelectPanel.
// That panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when that panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for that panel is:
A select widget sorted by name.
The widget needs to be rebuilt anytime a contact is added, edited or removed.

*/
func (group *panelGroup) showRemoveSelectPanel(force bool) {
	viewtools.ShowPanelInButtonGroup(group.removeSelectPanel, force)
}

// showRemoveFormPanel shows the panel you named RemoveFormPanel while hiding any other panels in this panel group.
// This panel's id is mainMasterView-home-pad-RemoveButton-RemoveFormPanel.
// This panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when this panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for this panel is:
A contact display with "remove" and "cancel" buttons.
*/
func (group *panelGroup) showRemoveFormPanel(force bool) {
	viewtools.ShowPanelInButtonGroup(group.removeFormPanel, force)
}
