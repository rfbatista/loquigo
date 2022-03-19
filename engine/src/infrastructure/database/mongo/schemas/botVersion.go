package schemas

import "loquigo/engine/src/core/domain"

func NewBotVersionSchema(botVersion domain.BotVersion) BotVersionSchema {
	return BotVersionSchema{ID: botVersion.ID, BotId: botVersion.BotId, Yaml: botVersion.Yaml}
}

type BotVersionSchema struct {
	ID    string `bson:"id"`
	BotId string `bson:"bot_id"`
	Yaml  string `bson:"yaml"`
}

func (b BotVersionSchema) ToDomain() domain.BotVersion {
	return domain.BotVersion{ID: b.ID, BotId: b.BotId, Yaml: b.Yaml}
}
