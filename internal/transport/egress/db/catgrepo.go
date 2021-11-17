package db

import (
	"context"

	"github.com/gimmickless/keat-kit-service/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const (
	catgCollName = "categories"
)

type CategoryRepository struct {
	logger *zap.SugaredLogger
	db     *mongo.Database
}

func NewCategoryRepository(logger *zap.SugaredLogger, db *mongo.Database) *CategoryRepository {
	return &CategoryRepository{logger, db}
}

func (r *CategoryRepository) Insert(ctx context.Context, catg domain.Category) (string, error) {
	return "", nil
}

func (r *CategoryRepository) Update(ctx context.Context, id string, catg domain.Category) error {
	return nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (r *CategoryRepository) Get(ctx context.Context, id string) (domain.Category, error) {
	return domain.Category{}, nil
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]domain.Category, error) {
	return nil, nil
}
