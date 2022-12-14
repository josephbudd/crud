package message

// ReloadContactsRendererToMainProcess is the ReloadContacts message that the renderer sends to the main process.
type ReloadContactsRendererToMainProcess struct {
}

// ReloadContactsMainProcessToRenderer is the ReloadContacts message that the main process sends to the renderer.
// Error indicates that an error occurred.
// ErrorMessage is the error message to show the user in the renderer process.
// Fatal indicates that the error is fatal and the renderer process must end.
// If Fatal is true then this message is automatically handled by the renderer process. You can ignore it.
type ReloadContactsMainProcessToRenderer struct {
	Error        bool
	ErrorMessage string
	Fatal        bool
}
