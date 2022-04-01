package mocks

import (
	"loquigo/engine/pkg/core/domain"

	"github.com/stretchr/testify/mock"
)

type MockBotRepository struct {
	mock.Mock
}

func (u MockBotRepository) Create(bot domain.Bot) (domain.Bot, error) {
	return domain.Bot{}, nil
}

func (u MockBotRepository) Update(bot domain.Bot) (domain.Bot, error) {
	return domain.Bot{}, nil
}

func (u MockBotRepository) Delete(bot domain.Bot) (domain.Bot, error) {
	return domain.Bot{}, nil
}

func (b MockBotRepository) FindById(botId string) (domain.Bot, error) {
	return domain.Bot{}, nil
}

func (b MockBotRepository) FindBeginByBotId(botId string) (string, error) {
	return "123", nil
}
func (b MockBotRepository) GetBots() ([]domain.Bot, error) {
	ret := b.Called()

	var r0 []domain.Bot
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]domain.Bot)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
