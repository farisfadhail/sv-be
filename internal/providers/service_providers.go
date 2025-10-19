//go:build wireinject

package providers

import (
	"test-be/internal/services"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	services.NewArticleService,
)
