package app

import (
	"context"

	"github.com/gimmickless/keat-kit-service/internal/domain"
	"github.com/gimmickless/keat-kit-service/pkg/enum"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type IIngredientService interface {
	Create(ctx context.Context, catg domain.Ingredient) (string, error)
	Update(ctx context.Context, id string, catg domain.Ingredient) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (domain.Ingredient, error)
	GetPage(
		ctx context.Context, page int, pageSize int, sortField string, sortDirection enum.SortDirection,
	) ([]domain.Ingredient, error)
	GetAll(ctx context.Context) ([]domain.Ingredient, error)
}

type IngredientService struct {
	logger     *otelzap.SugaredLogger
	ingredRepo IIngredientRepo
}

func NewIngredientService(
	logger *otelzap.SugaredLogger,
	ingredRepo IIngredientRepo,
) *IngredientService {
	return &IngredientService{logger, ingredRepo}
}

func (s *IngredientService) Create(ctx context.Context, catg domain.Ingredient) (string, error) {
	return s.ingredRepo.Insert(ctx, catg)
}

func (s *IngredientService) Update(ctx context.Context, id string, catg domain.Ingredient) error {
	return s.ingredRepo.Update(ctx, id, catg)
}

func (s *IngredientService) Delete(ctx context.Context, id string) error {
	return s.ingredRepo.Delete(ctx, id)
}

func (s *IngredientService) Get(ctx context.Context, id string) (domain.Ingredient, error) {
	return s.ingredRepo.Get(ctx, id)
}

func (s *IngredientService) GetPage(
	ctx context.Context, page int, pageSize int, sortField string, sortDirection enum.SortDirection,
) ([]domain.Ingredient, error) {
	return s.ingredRepo.GetPaginated(ctx, pageSize, page-1, sortField, sortDirection)
}

func (s *IngredientService) GetAll(ctx context.Context) ([]domain.Ingredient, error) {
	return s.ingredRepo.GetAll(ctx)
}
