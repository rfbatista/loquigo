//go:build wireinject
// +build wireinject

package main

import (
	"loquigo/engine/internal"

	//Infrastructure
	infra "loquigo/engine/pkg/infrastructure"

	"github.com/google/wire"
)

func InitializeEvent() (infra.Server, error) {
	wire.Build(internal.ChatProviderSet, infra.NewServer)
	return infra.Server{}, nil
}
