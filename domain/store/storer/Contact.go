package storer

import (
	"github.com/josephbudd/crud/domain/store/record"
)

// ContactStorer defines the behavior (API) of a store of /domain/store/record.Contact records.
type ContactStorer interface {

	// Open opens the data-store.
	// Returns the error.
	Open() (err error)

	// Close closes the data-store.
	// Returns the error.
	Close() (err error)

	// Get retrieves one *record.Contact from the data-store.
	// Param id is the record ID.
	// Returns a record pointer and error.
	// When no record is found, the returned record pointer is nil and the returned error is nil.
	Get(id uint64) (r *record.Contact, err error)

	// GetAll retrieves all of the *record.Contact records from the data-store.
	// Returns a slice of record pointers and error.
	// When no records are found, the returned slice length is 0 and the returned error is nil.
	GetAll() (rr []*record.Contact, err error)

	// Update updates the *record.Contact in the data-store.
	// Param r is the pointer to the record to be updated.
	// If r is a new record then r.ID is updated as well.
	// Returns the error.
	Update(r *record.Contact) (err error)

	// Remove removes the record.Contact from the data-store.
	// Param id is the record ID of the record to be removed.
	// Returns the error.
	Remove(id uint64) (err error)
}
