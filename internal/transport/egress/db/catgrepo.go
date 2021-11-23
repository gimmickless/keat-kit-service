package db

import (
	"context"
	"fmt"

	"github.com/gimmickless/keat-kit-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const (
	catgCollName = "categories"
)

type categoryFetchDAO struct {
	ObjectID primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Desc     string             `bson:"desc"`
	ImgPath  string             `bson:"imagePath"`
}

type CategoryRepository struct {
	logger   *zap.SugaredLogger
	catgColl *mongo.Collection
}

func NewCategoryRepository(logger *zap.SugaredLogger, db *mongo.Database) *CategoryRepository {
	return &CategoryRepository{logger, db.Collection(catgCollName)}
}

func (r *CategoryRepository) Insert(ctx context.Context, catg domain.Category) (string, error) {
	res, err := r.catgColl.InsertOne(ctx, bson.D{
		primitive.E{Key: "name", Value: catg.Name},
		primitive.E{Key: "desc", Value: catg.Desc},
		primitive.E{Key: "imgPath", Value: catg.ImgPath},
	})
	if err != nil {
		r.logger.Errorw("failed to insert category",
			"category", catg,
			"error", err,
		)
		return "", err
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return "", fmt.Errorf("failed to convert objectfor category: %v", catg.Name)
}

func (r *CategoryRepository) Update(ctx context.Context, id string, catg domain.Category) error {
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	res := r.catgColl.FindOneAndUpdate(
		ctx,
		bson.M{"_id": id},
		bson.D{
			primitive.E{Key: "name", Value: catg.Name},
			primitive.E{Key: "desc", Value: catg.Desc},
			primitive.E{Key: "imgPath", Value: catg.ImgPath},
		},
		&opt,
	)

	if err := res.Err(); err != nil {
		r.logger.Errorw("failed to update category",
			"category", catg,
			"error", err,
		)
		return err
	}
	return nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Errorw("could not convert id to objectID",
			"id", id,
			"error", err,
		)
	}

	res, err := r.catgColl.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		r.logger.Errorw("failed to delete category",
			"id", id,
			"error", err,
		)
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no campaign found to delete with id %s", id)
	}
	return nil
}

func (r *CategoryRepository) Get(ctx context.Context, id string) (domain.Category, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Errorw("could not convert id to objectID", "id", id, "error", err)
		return domain.Category{}, err
	}
	var catgDAO categoryFetchDAO
	if err = r.catgColl.FindOne(ctx, bson.M{"_id": objectID}).Decode(&catgDAO); err != nil {
		r.logger.Errorw("error getting category", "id", id, "error", err)
		return domain.Category{}, err
	}
	return convertToDomainCategory(catgDAO), nil
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]domain.Category, error) {
	var daos []categoryFetchDAO
	var opts = options.Find().SetSort(bson.D{
		{Key: "name", Value: 1},
	})
	cursor, err := r.catgColl.Find(
		ctx,
		bson.M{},
		opts,
	)
	if err != nil {
		r.logger.Errorw("failed to get categories from db", "error", err)
		return nil, err
	}
	if err = cursor.All(ctx, &daos); err != nil {
		r.logger.Errorw("failed to map category cursor to a list of DAO objects", "error", err)
		return nil, err
	}

	categories := make([]domain.Category, len(daos))
	for _, d := range daos {
		categories = append(categories, convertToDomainCategory(d))
	}
	return categories, nil
}

// Auxiliary functions
func convertToDomainCategory(dao categoryFetchDAO) domain.Category {
	return domain.Category{
		ID:      dao.ObjectID.Hex(),
		Name:    dao.Name,
		Desc:    dao.Desc,
		ImgPath: dao.ImgPath,
	}
}
