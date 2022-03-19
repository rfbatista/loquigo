package bot

import "loquigo/engine/pkg/core/domain"

func NewBotService(r BotRepository, v BotVersionRepository) BotService {
	return BotService{botrepo: r, versionrepo: v}
}

type BotService struct {
	botrepo     BotRepository
	versionrepo BotVersionRepository
}

func (b BotService) FindBotBegin(botId string) (string, error) {
	flowId, _ := b.botrepo.FindBeginByBotId(botId)
	return flowId, nil
}

func (b BotService) GetBots() ([]domain.Bot, error) {
	bots, _ := b.botrepo.GetBots()
	return bots, nil
}

//todo: implement
func (b BotService) GetBotVersions(botId string) (string, error) {
	flowId, _ := b.botrepo.FindBeginByBotId(botId)
	return flowId, nil
}

func (b BotService) CreateBot(bot domain.Bot) (domain.Bot, error) {
	saveBot, err := b.botrepo.Create(bot)
	if err != nil {
		return domain.Bot{}, err
	}
	return saveBot, nil
}

func (b BotService) DeleteBot(bot domain.Bot) (domain.Bot, error) {
	deletedBot, _ := b.botrepo.Delete(bot)
	return deletedBot, nil
}

func (b BotService) CreateVersion(bot domain.Bot, versionId string) (domain.BotVersion, error) {
	version := domain.BotVersion{ID: versionId, BotId: bot.ID}
	botVersionCreated, _ := b.versionrepo.Create(version)
	return botVersionCreated, nil
}

func (b BotService) FindVersionsByBotId(botId string) ([]domain.BotVersion, error) {
	versions, _ := b.versionrepo.FindByBotId(botId)
	return versions, nil
}
func (b BotService) FindVersionByIdAndBotId(versionId string, botId string) (domain.BotVersion, error) {
	version, _ := b.versionrepo.FindByIdAndBotId(versionId, botId)
	return version, nil
}
