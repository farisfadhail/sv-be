//go:build wireinject

package providers

import (
	"test-be/internal/handler"

	"github.com/google/wire"
)

var HandlerProviderSet = wire.NewSet(
	handler.NewArticleHandler,
)
