package storing

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/boltdb/bolt"
	"github.com/josephbudd/crud/domain/store/record"
)

/*
	type ContactBoltDB
	is the implementation of the /domain/store/storer.ContactStorer interface
	  for the bolt database.
*/

const contactBucketName string = "contact"

// ContactBoltDB is the bolt db for key codes.
type ContactBoltDB struct {
	DB    *bolt.DB
	path  string
	perms os.FileMode
}

// NewContactBoltDB constructs a new ContactBoltDB.
// Param db is an open bolt data-store.
// Returns a pointer to the new ContactBoltDB.
func NewContactBoltDB(db *bolt.DB, path string, perms os.FileMode) (store *ContactBoltDB) {
	store = &ContactBoltDB{
		DB:    db,
		path:  path,
		perms: perms,
	}
	return
}

// Open opens the bolt data-store.
// Returns the error.
func (store *ContactBoltDB) Open() (err error) {
	// the bolt db is already open
	return
}

// Close closes the bolt data-store.
// Returns the error.
func (store *ContactBoltDB) Close() (err error) {
	err = store.DB.Close()
	return
}

// Get retrieves the record.Contact from the bolt data-store.
// Param id is the record id.
// Returns the record and error.
// If no record is found returns a nil record and a nil error.
func (store *ContactBoltDB) Get(id uint64) (r *record.Contact, err error) {
	ids := fmt.Sprintf("%d", id)
	err = store.DB.View(func(tx *bolt.Tx) (er error) {
		bucketname := []byte(contactBucketName)
		var bucket *bolt.Bucket
		if bucket = tx.Bucket(bucketname); bucket == nil {
			return
		}
		var rbb []byte
		if rbb = bucket.Get([]byte(ids)); rbb == nil {
			// not found
			return
		}
		r = &record.Contact{}
		if er = json.Unmarshal(rbb, r); er != nil {
			r = nil
			return
		}
		return
	})
	return
}

// GetAll retrieves all of the record.Contact from the bolt data-store.
// If no record is found then it calls store.initialize() and tries again. See *ContactBoltDB.initialize().
// Returns the records and error.
// If no record is found returns a zero length records and a nil error.
func (store *ContactBoltDB) GetAll() (rr []*record.Contact, err error) {
	if rr, err = store.getAll(); len(rr) == 0 && err == nil {
		store.initialize()
		rr, err = store.getAll()
	}
	return
}

func (store *ContactBoltDB) getAll() (rr []*record.Contact, err error) {
	err = store.DB.View(func(tx *bolt.Tx) (er error) {
		bucketname := []byte(contactBucketName)
		var bucket *bolt.Bucket
		if bucket = tx.Bucket(bucketname); bucket == nil {
			return
		}
		c := bucket.Cursor()
		rr = make([]*record.Contact, 0, 1024)
		for k, v := c.First(); k != nil; k, v = c.Next() {
			r := record.NewContact()
			if er = json.Unmarshal(v, r); er != nil {
				rr = nil
				return
			}
			rr = append(rr, r)
		}
		return
	})
	return
}

// Update adds or updates a record.Contact in the bolt data-store.
// Param r is the record to be updated.
// If r is a new record then r.ID is updated with the new record id.
// Returns the error.
func (store *ContactBoltDB) Update(r *record.Contact) (err error) {
	err = store.update(r)
	return
}

// Remove removes a record.Contact from the bolt data-store.
// Param id the key of the record to be removed.
// If the record is not found returns a nil error.
// Returns the error.
func (store *ContactBoltDB) Remove(id uint64) (err error) {
	err = store.DB.Update(func(tx *bolt.Tx) (er error) {
		bucketname := []byte(contactBucketName)
		var bucket *bolt.Bucket
		if bucket = tx.Bucket(bucketname); bucket == nil {
			return
		}
		idbb := []byte(fmt.Sprintf("%d", id))
		er = bucket.Delete(idbb)
		return
	})
	return
}

// updates the record.Contact in the bolt data-store.
// Param record the record to be updated.
// If the record is new then it's ID is updated.
// Returns the error.
func (store *ContactBoltDB) update(r *record.Contact) (err error) {
	err = store.DB.Update(func(tx *bolt.Tx) (er error) {
		bucketname := []byte(contactBucketName)
		var bucket *bolt.Bucket
		if bucket, er = tx.CreateBucketIfNotExists(bucketname); er != nil {
			return
		}
		if r.ID == 0 {
			if r.ID, er = bucket.NextSequence(); er != nil {
				return
			}
		}
		var rbb []byte
		if rbb, er = json.Marshal(r); er != nil {
			return
		}
		idbb := []byte(fmt.Sprintf("%d", r.ID))
		er = bucket.Put(idbb, rbb)
		return
	})
	return
}

// initialize is only useful if you want to add the default records to the bolt data-store.
// otherwise you don't need it to do anything.
func (store *ContactBoltDB) initialize() (err error) {
	/*
		example code:

		defaults := somepackage.GetContactDefaults()
		for _, default := range defaults {
			r := types.NewContactRecord()
			r.Name = default.Name
			r.Price = default.Price
			r.SKU = default.SKU
			if err = store.update(r); err != nil {
				return
			}
		}
	*/
	return
}
