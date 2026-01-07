package campaign

import "time"

type Contact struct {
	Email string
}

type Campaign struct {
	Id 			string
	Name 		string
	Content		string
	CreatedAt 	time.Time
	Contacts 	[]Contact
}