package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

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

func NewCampaign(name, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("Name is required.")
	} else if content == "" {
		return nil, errors.New("Content is required.")
	} else if len(emails) == 0 {
		return nil, errors.New("E-mail is required.")
	}
	
	contacts := make([]Contact, len(emails))
	for index, value := range emails {
		contacts[index].Email = value
	}	

	return &Campaign{
		Id: xid.New().String(),
		Name: name,
		Content: content,	
		CreatedAt: time.Now(),
		Contacts: contacts,
	}, nil
}