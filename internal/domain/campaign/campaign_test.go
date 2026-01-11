package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	name 		= "Campaign X"
	content 	= "Body"
	contacts 	= []string{"example1@test.com", "example2@test.com"}
)

func TestNewCampaign(t *testing.T) {
	t.Run("should create new campaign", func(t *testing.T){
		assert := assert.New(t)
		require := require.New(t)
		now := time.Now().Add(-time.Minute)
		
		campaign, err := NewCampaign(name, content, contacts)

		require.NoError(err)
		assert.NotNil(campaign.Id)
		assert.Greater(campaign.CreatedAt, now)
		assert.Equal(name, campaign.Name)
		assert.Equal(content, campaign.Content)
		assert.Equal(len(contacts), len(campaign.Contacts))
	})

	t.Run("validation errors", func(t *testing.T){
		t.Run("should return error when name is empty", func(t *testing.T){
			assert := assert.New(t)
			require := require.New(t)

			_, err := NewCampaign("", content, contacts)

			require.Error(err)
			assert.EqualError(err, "Name is required.")
		})
		t.Run("should return error when content is empty", func(t *testing.T){
			assert := assert.New(t)
			require := require.New(t)

			_, err := NewCampaign(name, "", contacts)

			require.Error(err)
			assert.EqualError(err, "Content is required.")
		})
		t.Run("should return error when contacts is 0", func(t *testing.T){
			assert := assert.New(t)
			require := require.New(t)

			_, err := NewCampaign(name, content, []string{})

			require.Error(err)
			assert.EqualError(err, "E-mail is required.")
		})
	})
}