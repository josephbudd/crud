package message

import "github.com/josephbudd/crud/domain/store/record"

// GetRemoveSelectContactsPageRendererToMainProcess is the GetRemoveSelectContactsPage message that the renderer sends to the main process.
type GetRemoveSelectContactsPageRendererToMainProcess struct {
	FirstRecordIndex uint64
	Max              uint64
	State            uint64
}

// GetRemoveSelectContactsPageMainProcessToRenderer is the GetRemoveSelectContactsPage message that the main process sends to the renderer.
type GetRemoveSelectContactsPageMainProcessToRenderer struct {
	Error        bool
	ErrorMessage string

	Records          []*record.Contact
	FirstRecordIndex uint64
	State            uint64
	TotalRecordCount uint64
}
