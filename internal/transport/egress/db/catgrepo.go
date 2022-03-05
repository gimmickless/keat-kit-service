package db

import (
	"context"
	"fmt"

	"github.com/gimmickless/keat-kit-service/internal/domain"
	"github.com/gimmickless/keat-kit-service/pkg/custom"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	catgCollName = "categories"
)

type categoryFetchDAO struct {
	ObjectID primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Desc     string             `bson:"desc"`
	ImgPath  string             `bson:"imgPath"`
}

type CategoryRepository struct {
	logger   *otelzap.SugaredLogger
	catgColl *mongo.Collection
}

func NewCategoryRepository(logger *otelzap.SugaredLogger, db *mongo.Database) *CategoryRepository {
	return &CategoryRepository{logger, db.Collection(catgCollName)}
}

func (r *CategoryRepository) Insert(ctx context.Context, catg domain.Category) (string, error) {
	res, err := r.catgColl.InsertOne(ctx, bson.D{
		primitive.E{Key: "name", Value: catg.Name},
		primitive.E{Key: "desc", Value: catg.Desc},
		primitive.E{Key: "imgPath", Value: catg.ImgPath},
	})
	if err != nil {
		r.logger.Ctx(ctx).Errorw("failed to insert category", "category", catg, "error", err)
		return "", err
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return "", fmt.Errorf("failed to convert object for category %s", catg.Name)
}

func (r *CategoryRepository) Update(ctx context.Context, id string, catg domain.Category) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Ctx(ctx).Errorw("failed to convert string to mongo ObjectID", "id", id, "error", err)
		return err
	}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	res := r.catgColl.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": bson.D{
			primitive.E{Key: "name", Value: catg.Name},
			primitive.E{Key: "desc", Value: catg.Desc},
			primitive.E{Key: "imgPath", Value: catg.ImgPath},
		}},
		&opt,
	)

	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return &custom.ElemNotFoundError{ID: id, Err: fmt.Errorf("no category found to update")}
		}
		r.logger.Ctx(ctx).Errorw("failed to update category", "category", catg, "error", err)
		return err
	}
	return nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Ctx(ctx).Errorw("could not convert id to objectID", "id", id, "error", err)
	}

	res, err := r.catgColl.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		r.logger.Ctx(ctx).Errorw("failed to delete category", "id", id, "error", err)
		return err
	}
	if res.DeletedCount == 0 {
		return &custom.ElemNotFoundError{ID: id, Err: fmt.Errorf("no category found to delete")}
	}
	return nil
}

func (r *CategoryRepository) Get(ctx context.Context, id string) (domain.Category, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Ctx(ctx).Errorw("could not convert id to objectID", "id", id, "error", err)
		return domain.Category{}, err
	}
	var catgDAO categoryFetchDAO
	if err = r.catgColl.FindOne(ctx, bson.M{"_id": objectID}).Decode(&catgDAO); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Ctx(ctx).Debugw("no category found with id", "id", id)
			return domain.Category{}, &custom.ElemNotFoundError{ID: id, Err: err}
		}
		r.logger.Ctx(ctx).Errorw("error getting category", "id", id, "error", err)
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
		r.logger.Ctx(ctx).Errorw("failed to get categories from db", "error", err)
		return nil, err
	}
	if err = cursor.All(ctx, &daos); err != nil {
		r.logger.Ctx(ctx).Errorw("failed to map category cursor to a list of DAO objects", "error", err)
		return nil, err
	}

	catgs := make([]domain.Category, len(daos))
	for _, d := range daos {
		catgs = append(catgs, convertToDomainCategory(d))
	}
	return catgs, nil
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
