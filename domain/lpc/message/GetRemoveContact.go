package message

import "github.com/josephbudd/crud/domain/store/record"

// GetRemoveContactRendererToMainProcess is the GetRemoveContact message that the renderer sends to the main process.
type GetRemoveContactRendererToMainProcess struct {
	ID uint64
}

// GetRemoveContactMainProcessToRenderer is the GetRemoveContact message that the main process sends to the renderer.
type GetRemoveContactMainProcessToRenderer struct {
	Error        bool
	ErrorMessage string

	Record *record.Contact
	ID     uint64
}
