package campaign

import (
	"email_campaign/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (rm *repositoryMock) Save(campaign *Campaign) error {
	args := rm.Called(campaign)
	return args.Error(0)
}

func TestCreateCampaign(t *testing.T) {
	t.Run("success cases", func(t *testing.T){
		t.Run("should create a campaign and return (id, nil)", func(t *testing.T){
			assert := assert.New(t)
			newCampaign := contract.NewCampaign{
				Name: 		name,
				Content: 	content,
				Contacts: 	contacts,
			}
			rm := new(repositoryMock)
			rm.On("Save", mock.MatchedBy(func (c *Campaign) bool {
				return 	c.Name == newCampaign.Name &&
						c.Content == newCampaign.Content &&
						len(c.Contacts) == len(newCampaign.Contacts)
			})).Return(nil)
			service := Service{Repository: rm}
			
			id, err := service.Create(newCampaign)

			assert.NotNil(id)
			assert.Nil(err)
		})
		t.Run("should save a campaign and return nil", func(t *testing.T){
			newCampaign := contract.NewCampaign{
				Name:		name,
				Content: 	content,
				Contacts: 	contacts,
			}
			rm := new(repositoryMock)
			rm.On("Save", mock.MatchedBy(func (c *Campaign) bool {
				return 	c.Name == newCampaign.Name &&
						c.Content == newCampaign.Content &&
						len(c.Contacts) == len(newCampaign.Contacts)
			})).Return(nil)
			service := Service{Repository: rm}

			service.Create(newCampaign)
			
			rm.AssertExpectations(t)
		})
	})
}	