package viewtools

import (
	"strings"
	"syscall/js"
)

/*
	WARNING:

	DO NOT EDIT THIS FILE.

*/

// ForceTabButtonClick implements the behavior of a tab button being clicked by the user.
func (tools *Tools) ForceTabButtonClick(button js.Value) {
	tools.handleTabButtonOnClick(button)
}

func (tools *Tools) initializeTabBar() {
	notJS := tools.NotJS
	tools.tabberLastPanelID = ""
	tools.tabberTabBarLastPanel = make(map[string]string, 20)

	f := func(e Event) (nilReturn interface{}) {
		tools.handleTabButtonOnClick(e.Target)
		return
	}
	for id := range tools.tabberTabBarLastPanel {
		if len(id) > 0 {
			tabbar := notJS.GetElementByID(id)
			tools.AddEventHandler(f, tabbar, "click", false)
		}
	}
}

func (tools *Tools) handleTabButtonOnClick(button js.Value) {
	if !tools.HandleButtonClick() {
		return
	}
	tools.setTabButtonFocus(button)
	nextpanelid := tools.NotJS.ID(button) + "Panel"
	if nextpanelid != tools.tabberLastPanelID {
		// clear this level
		parts := strings.Split(nextpanelid, "-")
		nextpanellevel := parts[0]
		tools.IDHide(tools.tabberTabBarLastPanel[nextpanellevel])
		// show the next panel
		tools.IDShow(nextpanelid)
		// remember next panel. it is now the last panel.
		tools.tabberLastPanelID = nextpanelid
		tools.tabberTabBarLastPanel[nextpanellevel] = nextpanelid
	}
	tools.SizeApp()
}

func (tools *Tools) setTabButtonFocus(tabinfocus js.Value) {
	// focus the tab now in focus
	notJS := tools.NotJS
	notJS.ClassListReplaceClass(tabinfocus, UnSelectedTabClassName, SelectedTabClassName)
	p := notJS.ParentNode(tabinfocus)
	children := notJS.ChildrenSlice(p)
	for _, ch := range children {
		if ch != tabinfocus && notJS.TagName(ch) == "BUTTON" && notJS.ClassListContains(ch, SelectedTabClassName) {
			notJS.ClassListReplaceClass(ch, SelectedTabClassName, UnSelectedTabClassName)
		}
	}
}
