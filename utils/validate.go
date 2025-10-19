package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate = validator.New()

func ValidateAndBind[T any](c *fiber.Ctx) (*T, error, int) {
	method := strings.ToUpper(c.Method())

	if method != fiber.MethodPost && method != fiber.MethodPut {
		return nil, errors.New("method not allowed"), http.StatusMethodNotAllowed
	}

	var req T

	if err := c.BodyParser(&req); err != nil {
		fmt.Println("Error parsing request body:", err)
		return nil, errors.New("invalid request body"), http.StatusBadRequest
	}

	if err := Validate.Struct(req); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrs {
				return nil, fmt.Errorf("Field '%s' failed on '%s' validation", e.Field(), e.Tag()), http.StatusUnprocessableEntity
			}
		}
		return nil, err, http.StatusUnprocessableEntity
	}

	return &req, nil, 0
}
