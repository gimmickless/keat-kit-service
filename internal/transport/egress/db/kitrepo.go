package db

import (
	"context"
	"time"

	"github.com/gimmickless/keat-kit-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const (
	kitCollName = "kits"
)

type kitFetchDAO struct {
	ObjectID    primitive.ObjectID `bson:"_id"`
	CategoryIDs []string           `bson:"categoryIds"`
	Name        string             `bson:"name"`
	Version     string             `bson:"version"`
	Description string             `bson:"description"`
	Status      string             `bson:"status"`
	Author      string             `bson:"author"`
	Recipe      []struct {
		ingredientFetchDAO
		Quantity float64 `bson:"quantity"`
		Action   string  `bson:"action"`
	} `bson:"recipe"`
	Energy        float64   `bson:"energy"`
	Portions      float64   `bson:"portions"`
	PrepTime      int       `bson:"prepTime"`
	CreatedAt     time.Time `bson:"createdAt"`
	LastUpdatedAt time.Time `bson:"lastUpdatedAt"`
	Price         struct {
		Country  string `bson:"country"`
		Currency string `bson:"currency"`
	} `bson:"price"`
}

type KitRepository struct {
	logger  *zap.SugaredLogger
	kitColl *mongo.Collection
}

func NewKitRepository(logger *zap.SugaredLogger, db *mongo.Database) *KitRepository {
	return &KitRepository{logger, db.Collection(kitCollName)}
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

// Auxiliary functions
func convertToDomainKit(dao kitFetchDAO) domain.Kit {
	return domain.Kit{
		ID: dao.ObjectID.Hex(),
		// TODO: Convert other fields
	}
}
