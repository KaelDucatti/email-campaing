package campaign

import (
	"email_campaign/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, _ := NewCampaign(
		newCampaign.Name,
		newCampaign.Content,
		newCampaign.Contacts,
	)
	
	s.Repository.Save(campaign)
	
	return campaign.Id, nil
}