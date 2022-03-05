package http

import (
	"github.com/gimmickless/keat-kit-service/configs"
	"github.com/gimmickless/keat-kit-service/internal/app"
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type HTTPHandler struct {
	logger    *otelzap.SugaredLogger
	catgsrv   *app.CategoryService
	ingredsrv *app.IngredientService
	kitsrv    *app.KitService
	tracer    oteltrace.Tracer
}

// NewHTTPHandler constructs a new HTTPHandler.
func NewHTTPHandler(
	logger *otelzap.SugaredLogger,
	catgsrv *app.CategoryService,
	ingredsrv *app.IngredientService,
	kitsrv *app.KitService,
) *HTTPHandler {
	tracer := otel.Tracer(configs.App.OpenTelemetry.TracerName)
	return &HTTPHandler{logger, catgsrv, ingredsrv, kitsrv, tracer}
}

// Category handlers
func (h *HTTPHandler) GetCategories(c *fiber.Ctx) error {
	_, span := h.tracer.Start(c.Context(), "GetCategories")
	defer span.End()
	catgs, err := h.catgsrv.GetAll(c.Context())
	if err != nil {
		h.logger.Ctx(c.Context()).Errorw("Could not get categories", "err", err)
		return err
	}
	return c.JSON(catgs)
}

func (h *HTTPHandler) GetCategory(c *fiber.Ctx) error {
	catgID := c.Params("id")
	_, span := h.tracer.Start(c.Context(), "GetCategory", oteltrace.WithAttributes(attribute.String("id", catgID)))
	defer span.End()
	catg, err := h.catgsrv.Get(c.Context(), catgID)
	if err != nil {
		h.logger.Ctx(c.Context()).Errorw("Could not get category", "id", catgID, "err", err)
		return err
	}
	return c.JSON(catg)
}

func (h *HTTPHandler) SuggestCategory(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": "1",
	})
}

func (h *HTTPHandler) UpdateCategory(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(categoryResp{})
}

func (h *HTTPHandler) UploadCategoryImage(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(categoryResp{})
}

func (h *HTTPHandler) DeleteCategory(c *fiber.Ctx) error {
	catgID := c.Params("id")
	_, span := h.tracer.Start(c.Context(), "DeleteCategory", oteltrace.WithAttributes(attribute.String("id", catgID)))
	defer span.End()
	if !isAdmin(c) {
		h.logger.Ctx(c.Context()).Errorw("Not authorized to delete category", "id", catgID)
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	err := h.catgsrv.Delete(c.Context(), catgID)
	if err != nil {
		h.logger.Ctx(c.Context()).Errorw("Could not delete categories", "err", err)
		return err
	}
	return c.JSON(fiber.Map{
		"id": catgID,
	})
}

// Ingredient handlers
func (h *HTTPHandler) GetIngredients(c *fiber.Ctx) error {
	return c.JSON([]ingredientResp{})
}

func (h *HTTPHandler) GetIngredient(c *fiber.Ctx) error {
	return c.JSON(ingredientResp{})
}

func (h *HTTPHandler) SuggestIngredient(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": "1",
	})
}

func (h *HTTPHandler) UpdateIngredient(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(ingredientResp{})
}

func (h *HTTPHandler) UploadIngredientImage(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(categoryResp{})
}

func (h *HTTPHandler) DeleteIngredient(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

// Kit handlers
func (h *HTTPHandler) GetKits(c *fiber.Ctx) error {
	return c.JSON([]kitResp{})
}

func (h *HTTPHandler) GetKit(c *fiber.Ctx) error {
	return c.JSON(kitResp{})
}

func (h *HTTPHandler) CreateKit(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": "1",
	})
}

func (h *HTTPHandler) UpdateKit(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(kitResp{})
}

func (h *HTTPHandler) DeleteKit(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
