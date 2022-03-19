package bot

import "loquigo/engine/src/core/domain"

func NewBotService(r BotRepository) BotService {
	return BotService{repo: r}
}

type BotService struct {
	repo BotRepository
}

func (b BotService) FindBotBegin(botId string) (string, error) {
	flowId, _ := b.repo.FindBeginByBotId(botId)
	return flowId, nil
}

func (b BotService) GetBots() ([]domain.Bot, error) {
	bots, _ := b.repo.GetBots()
	return bots, nil
}

//todo: implement
func (b BotService) GetBotVersions(botId string) (string, error) {
	flowId, _ := b.repo.FindBeginByBotId(botId)
	return flowId, nil
}

func (b BotService) CreateBot(bot domain.Bot) (domain.Bot, error) {
	saveBot, err := b.repo.Create(bot)
	if err != nil {
		return domain.Bot{}, err
	}
	return saveBot, nil
}

func (b BotService) DeleteBot(bot domain.Bot) (domain.Bot, error) {
	deletedBot, _ := b.repo.Delete(bot)
	return deletedBot, nil
}
