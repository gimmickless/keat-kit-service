package main

import (
	inhttp "github.com/gimmickless/keat-kit-service/internal/transport/ingress/http"
	"github.com/gimmickless/keat-kit-service/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func initRoutes(app *fiber.App, handler *inhttp.HTTPHandler) {
	v1api := app.Group("/api/v1")
	v1catgs := v1api.Group("/categories")
	v1ingreds := v1api.Group("/ingredients")
	v1kits := v1api.Group("/kits")

	v1adm := v1api.Group("/admin", middleware.OnlyAdmin)
	v1admcatgs := v1adm.Group("/categories")
	v1admingreds := v1adm.Group("/ingredients")
	v1admkits := v1adm.Group("/kits")

	// User Ops
	v1catgs.Get("/", handler.GetCategories)
	v1catgs.Get("/:id", handler.GetCategory)
	v1catgs.Post("/", handler.SuggestCategory)

	v1ingreds.Get("/", handler.GetIngredients)
	v1ingreds.Get("/:id", handler.GetIngredient)
	v1ingreds.Post("/", handler.SuggestIngredient)

	v1kits.Get("/", handler.GetKits)
	v1kits.Get("/:id", handler.GetKit)
	v1kits.Post("/", handler.CreateKit)
	v1kits.Put("/:id", handler.UpdateKit)
	v1kits.Delete("/:id", handler.DeleteKit)

	// Admin Ops
	v1admcatgs.Put("/:id", handler.UpdateCategory)
	v1admcatgs.Delete("/:id", handler.DeleteCategory)
	v1admcatgs.Post("/upload", handler.UploadCategoryImage)

	v1admingreds.Put("/:id", handler.UpdateIngredient)
	v1admingreds.Delete("/:id", handler.DeleteIngredient)
	v1admcatgs.Post("/upload", handler.UploadIngredientImage)

	v1admkits.Put("/:id", handler.UpdateKit)
	v1admkits.Delete("/:id", handler.DeleteKit)
}
