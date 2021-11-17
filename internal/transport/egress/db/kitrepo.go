package db

import (
	"context"

	"github.com/gimmickless/keat-kit-service/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const (
	kitCollName = "kits"
)

type KitRepository struct {
	logger *zap.SugaredLogger
	db     *mongo.Database
}

func NewKitRepository(logger *zap.SugaredLogger, db *mongo.Database) *KitRepository {
	return &KitRepository{logger, db}
}

func (r *KitRepository) Insert(ctx context.Context, catg domain.Kit) (string, error) {
	return "", nil
}

func (r *KitRepository) Update(ctx context.Context, id string, catg domain.Kit) error {
	return nil
}

func (r *KitRepository) UpdatePrice(ctx context.Context, id string, price domain.Price) error {
	return nil
}

func (r *KitRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (r *KitRepository) Get(ctx context.Context, id string) (domain.Kit, error) {
	return domain.Kit{}, nil
}

func (r *KitRepository) GetAll(ctx context.Context) ([]domain.Kit, error) {
	return nil, nil
}
