//go:build wireinject
// +build wireinject

package cmd

import (
	infra "loquigo/engine/src/infrastructure"
	"loquigo/engine/src/infrastructure/database/mongo"

	"github.com/google/wire"
)

func InitializeEvent(db mongo.MongoDB) (infra.Server, error) {
	wire.Build(
		ChatAndEditorServiceSet,
		infra.NewServer)
	return infra.Server{}, nil
}
