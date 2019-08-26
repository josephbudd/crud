package message

import "github.com/josephbudd/crud/domain/store/record"

// GetEditContactRendererToMainProcess is the GetEditContact message that the renderer sends to the main process.
type GetEditContactRendererToMainProcess struct {
	ID uint64
}

// GetEditContactMainProcessToRenderer is the GetEditContact message that the main process sends to the renderer.
type GetEditContactMainProcessToRenderer struct {
	Error        bool
	ErrorMessage string

	Record *record.Contact
	ID     uint64
}
