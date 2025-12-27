package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status       string      `json:"status"`
	ResponseCode int         `json:"response_code"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Error        string      `json:"error,omitempty"`
}

func Success(c *fiber.Ctx, message string, data interface{}) error {
	response := Response{
		Status:       "success",
		ResponseCode: fiber.StatusOK,
		Message:      message,
		Data:         data,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func Created(c *fiber.Ctx, message string, data interface{}) error {
	response := Response{
		Status:       "created",
		ResponseCode: fiber.StatusCreated,
		Message:      message,
		Data:         data,
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

func BadRequest(c *fiber.Ctx, message string, err string) error {
	response := Response{
		Status:       "error bad request",
		ResponseCode: fiber.StatusBadRequest,
		Message:      message,
		Error:        err,
	}
	return c.Status(fiber.StatusBadRequest).JSON(response)
}

func NotFound(c *fiber.Ctx, message string, err string) error {
	response := Response{
		Status:       "error not found",
		ResponseCode: fiber.StatusNotFound,
		Message:      message,
		Error:        err,
	}
	return c.Status(fiber.StatusNotFound).JSON(response)
}

func Unauthorized(c *fiber.Ctx, message string, err string) error {
	response := Response{
		Status:       "error unauthorized",
		ResponseCode: fiber.StatusUnauthorized,
		Message:      message,
		Error:        err,
	}
	return c.Status(fiber.StatusUnauthorized).JSON(response)
}
