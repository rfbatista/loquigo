package bot

import (
	"loquigo/engine/src/core/domain"
	"loquigo/engine/src/infrastructure/database/mongo/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBotService(t *testing.T) {
	mockedRepo := new(mocks.MockBotRepository)
	service := BotService{repo: mockedRepo}
	t.Run("it get all bots created", func(t *testing.T) {
		t.Run("it call GetBots in repository", func(t *testing.T) {
			expect := []domain.Bot{
				{ID: "123", BeginId: "TEST"},
			}
			mockedRepo.On("GetBots").Return(expect, nil)
			bots, err := service.GetBots()

			assert.NoError(t, err)
			assert.Equal(t, bots, expect)
			mockedRepo.AssertExpectations(t)
		})
	})
	t.Run("it get all versions by bot id", func(t *testing.T) {})
}
