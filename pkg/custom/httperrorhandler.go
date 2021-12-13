package custom

import "github.com/gofiber/fiber/v2"

func CreateCustomHTTPErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError
		message := ""

		// Retrieve the custom status code if it's an fiber.*Error
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			message = err.Error()
		}
		if _, ok := err.(*ElemNotFoundError); ok {
			code = fiber.StatusNotFound
			message = err.Error()
		}
		return ctx.Status(code).SendString(message)
	}
}
