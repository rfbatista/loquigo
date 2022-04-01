//go:build wireinject
// +build wireinject

package cmd

import (
	"go.uber.org/zap"
	"loquigo/engine/pkg/infrastructure"
	"loquigo/engine/pkg/infrastructure/database/mongo"

	"github.com/google/wire"
)

func InitializeEvent(db mongo.MongoDB, logger *zap.Logger) (infrastructure.Server, error) {
	wire.Build(
		ControllersSet,
		infrastructure.NewServer)
	return infrastructure.Server{}, nil
}
