package bot

import "loquigo/engine/src/core/domain"

type BotRepository interface {
	FindBeginByBotId(botId string) (string, error)
	GetBots() ([]domain.Bot, error)
	Create(bot domain.Bot) (domain.Bot, error)
	Delete(bot domain.Bot) (domain.Bot, error)
}
