//go:build wireinject

package providers

import (
	"test-be/internal/repositories"

	"github.com/google/wire"
)

var RepositoryProviderSet = wire.NewSet(
	repositories.NewArticleRepository,
)
