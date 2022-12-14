package message

import "github.com/josephbudd/crud/domain/store/record"

// EditContactRendererToMainProcess is the EditContact message that the renderer sends to the main process.
type EditContactRendererToMainProcess struct {
	Record *record.Contact
}

// EditContactMainProcessToRenderer is the EditContact message that the main process sends to the renderer.
// Error indicates that an error occurred.
// ErrorMessage is the error message to show the user in the renderer process.
// Fatal indicates that the error is fatal and the renderer process must end.
// If Fatal is true then this message is automatically handled by the renderer process. You can ignore it.
type EditContactMainProcessToRenderer struct {
	Record *record.Contact

	Error        bool
	ErrorMessage string
	Fatal        bool
}
