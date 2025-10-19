package routes

import (
	"test-be/internal/injector"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupRouter(app *fiber.App, ct *injector.AppContainer) {
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Article",
		})
	})

	api := app.Group("/api")

	tx := api.Group("/article")
	{
		tx.Get("/", ct.ArticleHandler.Index)
		tx.Post("/", ct.ArticleHandler.Create)
		tx.Put("/:id", ct.ArticleHandler.Update)
		tx.Get("/:id", ct.ArticleHandler.GetByID)
		tx.Delete("/:id", ct.ArticleHandler.Delete)
	}

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Endpoint not found",
		})
	})
}
