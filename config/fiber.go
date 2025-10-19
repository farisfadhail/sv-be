package config

import (
	"github.com/gofiber/fiber/v2"
)

func NewFiber() *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      GetEnv("APP_NAME", "SV-BE"),
		ErrorHandler: NewErrorHandler(),
		Prefork:      GetEnv("APP_PREFORK", "false") == "true",
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
