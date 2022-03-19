//go:build wireinject
// +build wireinject

package cmd

import (
	infra "loquigo/engine/pkg/infrastructure"
	"loquigo/engine/pkg/infrastructure/database/mongo"

	"github.com/google/wire"
)

func InitializeEvent(db mongo.MongoDB) (infra.Server, error) {
	wire.Build(
		ControllersSet,
		infra.NewServer)
	return infra.Server{}, nil
}
