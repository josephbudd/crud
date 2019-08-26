package record

/*

	TODO:

	You need to complete this local bolt store record definition.

*/

// Contact is the local bolt store Contact record.
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

// NewContact constructs a new local bolt store Contact.
func NewContact() *Contact {
	v := &Contact{}
	return v
}
