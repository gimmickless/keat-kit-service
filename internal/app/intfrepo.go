package app

import (
	"context"

	"github.com/gimmickless/keat-kit-service/internal/domain"
)

type ICategoryRepo interface {
	Insert(ctx context.Context, catg domain.Category) (string, error)
	Update(ctx context.Context, id string, catg domain.Category) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (domain.Category, error)
	GetAll(ctx context.Context) ([]domain.Category, error)
}

type IIngredientRepo interface {
	Insert(ctx context.Context, catg domain.Ingredient) (string, error)
	Update(ctx context.Context, id string, catg domain.Ingredient) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (domain.Ingredient, error)
	GetAll(ctx context.Context) ([]domain.Ingredient, error)
}

type IKitRepo interface {
	Insert(ctx context.Context, catg domain.Kit) (string, error)
	Update(ctx context.Context, id string, catg domain.Kit) error
	UpdatePrice(ctx context.Context, id string, price domain.Price) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (domain.Kit, error)
	GetAll(ctx context.Context) ([]domain.Kit, error)
}
