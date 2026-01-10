package campaign

import (
	"testing"
)

func TestNewCampaing(t *testing.T) {
	name := "Campaign X"
	content := "Body"
	contacts := []string{"example1@test.com", "example2@test.com"}
	campaign := NewCampaign(name, content, contacts)

	if campaign.Id != 1 {
		t.Errorf("Expected: 1")
	} else if campaign.Name != name {
		t.Errorf("Expected: correct name\n")
	} else if campaign.Content != content {
		t.Errorf("Expected correct content\n")
	} else if len(campaign.Contacts) != len(contacts) {
		t.Errorf("Expected correct content\n")
	}
}