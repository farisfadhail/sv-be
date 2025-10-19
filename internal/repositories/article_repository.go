package repositories

import (
	"test-be/models"

	"gorm.io/gorm"
)

type ArticleRepository struct {
}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{}
}

func (r *ArticleRepository) Create(db *gorm.DB, article models.Article) error {
	return db.Create(&article).Error
}

func (r *ArticleRepository) GetByID(db *gorm.DB, id string, preload ...string) (*models.Article, error) {
	var article *models.Article

	query := db.Model(&article).Where("id = ?", id)

	for _, p := range preload {
		query = query.Preload(p)
	}

	err := query.First(&article).Error

	return article, err
}

func (r *ArticleRepository) Update(db *gorm.DB, article *models.Article) error {
	return db.Save(article).Error
}
