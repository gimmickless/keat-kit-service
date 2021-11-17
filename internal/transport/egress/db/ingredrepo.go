package db

import (
	"context"

	"github.com/gimmickless/keat-kit-service/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const (
	ingredCollName = "ingredients"
)

type IngredientRepository struct {
	logger *zap.SugaredLogger
	db     *mongo.Database
}

func NewIngredientRepository(logger *zap.SugaredLogger, db *mongo.Database) *IngredientRepository {
	return &IngredientRepository{logger, db}
}

func (r *IngredientRepository) Insert(ctx context.Context, catg domain.Ingredient) (string, error) {
	return "", nil
}

func (r *IngredientRepository) Update(ctx context.Context, id string, catg domain.Ingredient) error {
	return nil
}

func (r *IngredientRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (r *IngredientRepository) Get(ctx context.Context, id string) (domain.Ingredient, error) {
	return domain.Ingredient{}, nil
}

func (r *IngredientRepository) GetAll(ctx context.Context) ([]domain.Ingredient, error) {
	return nil, nil
}
