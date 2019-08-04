package message

import "github.com/josephbudd/crud/domain/store/record"

// GetEditSelectContactsPageRendererToMainProcess is the GetEditSelectContactsPage message that the renderer sends to the main process.
type GetEditSelectContactsPageRendererToMainProcess struct {
	FirstRecordIndex uint64
	Max              uint64
	State            uint64
}

// GetEditSelectContactsPageMainProcessToRenderer is the GetEditSelectContactsPage message that the main process sends to the renderer.
type GetEditSelectContactsPageMainProcessToRenderer struct {
	Error        bool
	ErrorMessage string

	Records          []*record.Contact
	FirstRecordIndex uint64
	State            uint64
	TotalRecordCount uint64
}
