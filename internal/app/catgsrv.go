package app

import (
	"context"

	"github.com/gimmickless/keat-kit-service/internal/domain"
	"go.uber.org/zap"
)

type ICategoryService interface {
	Create(ctx context.Context, catg domain.Category) (string, error)
	Update(ctx context.Context, id string, catg domain.Category) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (domain.Category, error)
	GetAll(ctx context.Context) ([]domain.Category, error)
}

type CategoryService struct {
	logger   *zap.SugaredLogger
	catgRepo ICategoryRepo
}

func NewCategoryService(
	logger *zap.SugaredLogger,
	catgRepo ICategoryRepo,
) *CategoryService {
	return &CategoryService{logger, catgRepo}
}

func (s *CategoryService) Create(ctx context.Context, catg domain.Category) (string, error) {
	return s.catgRepo.Insert(ctx, catg)
}

func (s *CategoryService) Update(ctx context.Context, id string, catg domain.Category) error {
	return s.catgRepo.Update(ctx, id, catg)
}

func (s *CategoryService) Delete(ctx context.Context, id string) error {
	return s.catgRepo.Delete(ctx, id)
}

func (s *CategoryService) Get(ctx context.Context, id string) (domain.Category, error) {
	return s.catgRepo.Get(ctx, id)
}

func (s *CategoryService) GetAll(ctx context.Context) ([]domain.Category, error) {
	return s.catgRepo.GetAll(ctx)
}
