package resources

import "test-be/models"

type ArticleResource struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

func ToArticleResource(article models.Article) *ArticleResource {
	return &ArticleResource{
		Title:    article.Title,
		Content:  article.Content,
		Category: article.Category,
		Status:   article.Status,
	}
}
