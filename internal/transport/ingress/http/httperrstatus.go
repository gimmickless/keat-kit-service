package http

import "github.com/gofiber/fiber/v2"

func getHTTPStatus(err error) *fiber.Error {

	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
}
