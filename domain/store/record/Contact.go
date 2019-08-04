package record

/*

	TODO:

	You need to complete this record definition.

*/

// Contact is a Contact record.
type Contact struct {
	ID       uint64
	Name     string
	Address1 string
	Address2 string
	City     string
	State    string
	Zip      string
	Phone    string
	Email    string
	Social   string
}

// NewContact constructs a new Contact record.
func NewContact() *Contact {
	v := &Contact{}
	return v
}
