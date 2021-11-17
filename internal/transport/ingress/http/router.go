package http

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, handler *HTTPHandler) {
	v1api := app.Group("/api/v1")
	v1catgs := v1api.Group("/categories")
	v1ingreds := v1api.Group("/ingredients")
	v1kits := v1api.Group("/kits")

	v1catgs.Get("/", handler.GetCategories)
	v1catgs.Get("/:id", handler.GetCategory)
	v1catgs.Post("/", handler.CreateCategory)
	v1catgs.Put("/:id", handler.UpdateCategory)
	v1catgs.Delete("/:id", handler.DeleteCategory)

	v1ingreds.Get("/", handler.GetIngredients)
	v1ingreds.Get("/:id", handler.GetIngredient)
	v1ingreds.Post("/", handler.CreateIngredient)
	v1ingreds.Put("/:id", handler.UpdateIngredient)
	v1ingreds.Delete("/:id", handler.DeleteIngredient)

	v1kits.Get("/", handler.GetKits)
	v1kits.Get("/:id", handler.GetKit)
	v1kits.Post("/", handler.CreateKit)
	v1kits.Put("/:id", handler.UpdateKit)
	v1kits.Delete("/:id", handler.DeleteKit)
}
