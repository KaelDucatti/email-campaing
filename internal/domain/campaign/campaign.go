package campaign

import "time"

type Contact struct {
	Email string
}

type Campaign struct {
	Id 			int
	Name 		string
	Content		string
	CreatedAt 	time.Time
	Contacts 	[]Contact
}

func NewCampaign(name, content string, emails []string) *Campaign {
	contacts := make([]Contact, len(emails))
	for index, value := range emails {
		contacts[index].Email = value
	}	

	return &Campaign{
		Id: 1,
		Name: name,
		Content: content,
		CreatedAt: time.Now(),
		Contacts: contacts,
	}
}