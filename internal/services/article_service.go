package services

import (
	"strconv"
	"test-be/internal/repositories"
	"test-be/internal/validations"
	"test-be/models"
	"test-be/resources"

	"gorm.io/gorm"
)

type ArticleService struct {
	db          *gorm.DB
	articleRepo *repositories.ArticleRepository
}

func NewArticleService(
	db *gorm.DB,
	articleRepo *repositories.ArticleRepository,
) *ArticleService {
	return &ArticleService{
		db:          db,
		articleRepo: articleRepo,
	}
}

func (s *ArticleService) FindAll(limit, offset, status string) ([]resources.ArticleResource, error) {
	var articles []models.Article

	query := s.db

	lim, err := strconv.Atoi(limit)
	if err != nil || lim <= 0 {
		lim = 10
	}
	off, err := strconv.Atoi(offset)
	if err != nil || off < 0 {
		off = 0
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	err = query.Limit(lim).Offset(off).Find(&articles).Error
	if err != nil {
		return nil, err
	}

	var articleResources []resources.ArticleResource

	for _, article := range articles {
		articleResources = append(articleResources, *resources.ToArticleResource(article))
	}

	return articleResources, nil
}

func (s *ArticleService) Create(req validations.CreateArticleRequest) error {
	article := models.Article{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Status:   req.Status,
	}

	return s.articleRepo.Create(s.db, article)
}

func (s *ArticleService) Update(id string, req validations.UpdateArticleRequest) error {
	article, err := s.articleRepo.GetByID(s.db, id)
	if err != nil {
		return err
	}

	if req.Title != "" {
		article.Title = req.Title
	}

	if req.Content != "" {
		article.Content = req.Content
	}

	if req.Category != "" {
		article.Category = req.Category
	}

	if req.Status != "" {
		article.Status = req.Status
	}

	return s.articleRepo.Update(s.db, article)
}

func (s *ArticleService) GetByID(id string) (*resources.ArticleResource, error) {
	article, err := s.articleRepo.GetByID(s.db, id)
	if err != nil {
		return nil, err
	}

	return resources.ToArticleResource(*article), nil
}

func (s *ArticleService) Delete(id string) error {
	article, err := s.articleRepo.GetByID(s.db, id)
	if err != nil {
		return err
	}

	return s.db.Delete(&article).Error
}
