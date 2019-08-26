package message

import "github.com/josephbudd/crud/domain/store/record"

// AddContactRendererToMainProcess is the AddContact message that the renderer sends to the main process.
type AddContactRendererToMainProcess struct {
	Record *record.Contact
}

// AddContactMainProcessToRenderer is the AddContact message that the main process sends to the renderer.
type AddContactMainProcessToRenderer struct {
	Error        bool
	ErrorMessage string

	Record *record.Contact
}
