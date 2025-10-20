package handler

import (
	"test-be/internal/services"
	"test-be/internal/validations"
	"test-be/utils"

	"github.com/gofiber/fiber/v2"
)

type ArticleHandler struct {
	s *services.ArticleService
}

func NewArticleHandler(s *services.ArticleService) *ArticleHandler {
	return &ArticleHandler{s: s}
}

func (h *ArticleHandler) Index(c *fiber.Ctx) error {
	limit := c.Query("limit", "10")
	offset := c.Query("offset", "0")
	status := c.Query("status", "")

	articleResources, err := h.s.FindAll(limit, offset, status)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(articleResources)
}

func (h *ArticleHandler) Create(c *fiber.Ctx) error {
	req, err, code := utils.ValidateAndBind[validations.CreateArticleRequest](c)
	if err != nil {
		return c.Status(code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.s.Create(*req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Article created successfully",
	})
}

func (h *ArticleHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	req, err, code := utils.ValidateAndBind[validations.UpdateArticleRequest](c)
	if err != nil {
		return c.Status(code).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.s.Update(id, *req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Article updated successfully",
	})
}

func (h *ArticleHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	articleResource, err := h.s.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(articleResource)
}

func (h *ArticleHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.s.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Article deleted successfully",
	})
}
