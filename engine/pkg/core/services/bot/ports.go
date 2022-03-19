package bot

import "loquigo/engine/pkg/core/domain"

type BotRepository interface {
	FindBeginByBotId(botId string) (string, error)
	GetBots() ([]domain.Bot, error)
	Create(bot domain.Bot) (domain.Bot, error)
	Delete(bot domain.Bot) (domain.Bot, error)
}

type BotVersionRepository interface {
	Create(bot domain.BotVersion) (domain.BotVersion, error)
	FindByBotId(botId string) ([]domain.BotVersion, error)
	FindByIdAndBotId(versionId string, botId string) (domain.BotVersion, error)
}
