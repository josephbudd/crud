package message

// RemoveContactRendererToMainProcess is the RemoveContact message that the renderer sends to the main process.
type RemoveContactRendererToMainProcess struct {
	ID uint64
}

// RemoveContactMainProcessToRenderer is the RemoveContact message that the main process sends to the renderer.
type RemoveContactMainProcessToRenderer struct {
	Error        bool
	ErrorMessage string
	ID           uint64
}
