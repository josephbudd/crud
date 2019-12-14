// +build js, wasm

package printnotreadypanel

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
	printNotReadyPanel js.Value
	printSelectPanel js.Value
	printPrintPanel js.Value
}

func (group *panelGroup) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(group *panelGroup) defineMembers()")
		}
	}()

    var panel *markup.Element
 if panel = document.ElementByID("mainMasterView-home-pad-PrintButton-PrintNotReadyPanel"); panel == nil {
		err = errors.New("unable to find #mainMasterView-home-pad-PrintButton-PrintNotReadyPanel")
		return
    }
    group.printNotReadyPanel = panel.JSValue()
 if panel = document.ElementByID("mainMasterView-home-pad-PrintButton-PrintSelectPanel"); panel == nil {
		err = errors.New("unable to find #mainMasterView-home-pad-PrintButton-PrintSelectPanel")
		return
    }
    group.printSelectPanel = panel.JSValue()
 if panel = document.ElementByID("mainMasterView-home-pad-PrintButton-PrintPrintPanel"); panel == nil {
		err = errors.New("unable to find #mainMasterView-home-pad-PrintButton-PrintPrintPanel")
		return
    }
    group.printPrintPanel = panel.JSValue()

	return
}

/*
	Show panel funcs.

	Call these from the controller, presenter and messenger.
*/

// showPrintNotReadyPanel shows the panel you named PrintNotReadyPanel while hiding any other panels in this panel group.
// This panel's id is mainMasterView-home-pad-PrintButton-PrintNotReadyPanel.
// This panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when this panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for this panel is:
Static text displayed when there are no contacts in the store.
*/
func (group *panelGroup) showPrintNotReadyPanel(force bool) {
	viewtools.ShowPanelInButtonGroup(group.printNotReadyPanel, force)
}

// showPrintSelectPanel shows the panel you named PrintSelectPanel while hiding any other panels in this panel group.
// That panel's id is mainMasterView-home-pad-PrintButton-PrintSelectPanel.
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
func (group *panelGroup) showPrintSelectPanel(force bool) {
	viewtools.ShowPanelInButtonGroup(group.printSelectPanel, force)
}

// showPrintPrintPanel shows the panel you named PrintPrintPanel while hiding any other panels in this panel group.
// That panel's id is mainMasterView-home-pad-PrintButton-PrintPrintPanel.
// That panel either becomes visible immediately or whenever this group of panels is made visible.  Whenever could be immediately if this panel group is currently visible.
// Param force boolean effects when that panel becomes visible.
//  * if force is true then
//    immediately if the home button pad is not currently displayed;
//    whenever if the home button pad is currently displayed.
//  * if force is false then whenever.
/* Your note for that panel is:
Display a record for printing.
*/
func (group *panelGroup) showPrintPrintPanel(force bool) {
	viewtools.ShowPanelInButtonGroup(group.printPrintPanel, force)
}
