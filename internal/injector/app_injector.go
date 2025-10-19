//go:build wireinject

//go:generate wire

package injector

import (
	"test-be/internal/handler"
	"test-be/internal/providers"

	"github.com/google/wire"
	"gorm.io/gorm"
)

type AppContainer struct {
	DB *gorm.DB
	//
	ArticleHandler *handler.ArticleHandler
}

func NewAppContainer(
	db *gorm.DB,
	//
	articleHandler *handler.ArticleHandler,
) (*AppContainer, error) {
	return &AppContainer{
		DB: db,
		//
		ArticleHandler: articleHandler,
	}, nil
}

func InitializeApp() (*AppContainer, error) {
	wire.Build(
		providers.AppProviderSet,
		NewAppContainer,
	)
	return nil, nil
}
