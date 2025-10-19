//go:build wireinject

package providers

import (
	"test-be/config"

	"github.com/google/wire"
)

var DatabaseProviderSet = wire.NewSet(
	config.ConnectGormDB,
)
