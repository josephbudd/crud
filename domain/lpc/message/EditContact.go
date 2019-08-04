package message

import (
	"github.com/josephbudd/crud/domain/store/record"
)

// EditContactRendererToMainProcess is the EditContact message that the renderer sends to the main process.
type EditContactRendererToMainProcess struct {
	Record *record.Contact
}

// EditContactMainProcessToRenderer is the EditContact message that the main process sends to the renderer.
type EditContactMainProcessToRenderer struct {
	Error        bool
	ErrorMessage string
	Record       *record.Contact
}
