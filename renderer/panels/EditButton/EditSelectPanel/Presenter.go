package editselectpanel

import (
	"github.com/pkg/errors"
)

/*

	Panel name: EditSelectPanel

*/

// panelPresenter writes to the panel
type panelPresenter struct {
	group     *panelGroup
	controler *panelControler
	caller    *panelCaller

	/* NOTE TO DEVELOPER: Step 1 of 3.

	// Declare your panelPresenter members here.

	*/
}

// defineMembers defines the panelPresenter members by their html elements.
// Returns the error.
func (presenter *panelPresenter) defineMembers() (err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "(presenter *panelPresenter) defineMembers()")
		}
	}()

	/* NOTE TO DEVELOPER. Step 2 of 3.

	// Define your panelPresenter members.

	*/

	return
}

/* NOTE TO DEVELOPER. Step 3 of 3.

// Define your panelPresenter functions.

*/
