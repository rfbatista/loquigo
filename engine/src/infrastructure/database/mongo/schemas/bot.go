package schemas

import "loquigo/engine/src/core/domain"

func NewBotSchema(bot domain.Bot) BotSchema {
	return BotSchema{Id: bot.ID, BeginId: bot.BeginId, Name: bot.Name, Version: bot.Version}
}

type BotSchema struct {
	Id      string `bson:"id"`
	BeginId string `bson:"begin_id"`
	Name    string `bson:"name"`
	Version string `bson:"version"`
}

func (b BotSchema) ToDomain() domain.Bot {
	return domain.Bot{ID: b.Id, BeginId: b.BeginId, Name: b.Name, Version: b.Version}
}
