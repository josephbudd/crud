package message

import "github.com/josephbudd/crud/domain/store/record"

// GetPrintSelectContactsPageRendererToMainProcess is the GetPrintSelectContactsPage message that the renderer sends to the main process.
type GetPrintSelectContactsPageRendererToMainProcess struct {
	FirstRecordIndex uint64
	Max              uint64
	State            uint64
}

// GetPrintSelectContactsPageMainProcessToRenderer is the GetPrintSelectContactsPage message that the main process sends to the renderer.
// Error indicates that an error occurred.
// ErrorMessage is the error message to show the user in the renderer process.
// Fatal indicates that the error is fatal and the renderer process must end.
// If Fatal is true then this message is automatically handled by the renderer process. You can ignore it.
type GetPrintSelectContactsPageMainProcessToRenderer struct {
	Records          []*record.Contact
	FirstRecordIndex uint64
	State            uint64
	TotalRecordCount uint64

	Error        bool
	ErrorMessage string
	Fatal        bool
}
