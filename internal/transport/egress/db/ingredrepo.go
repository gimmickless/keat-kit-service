package db

import (
	"context"
	"fmt"

	"github.com/gimmickless/keat-kit-service/internal/domain"
	"github.com/gimmickless/keat-kit-service/pkg/custom"
	"github.com/gimmickless/keat-kit-service/pkg/enum"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const (
	ingredCollName = "ingredients"
)

type ingredientFetchDAO struct {
	ObjectID   primitive.ObjectID `bson:"_id"`
	Code       string             `bson:"code"`
	Name       string             `bson:"name"`
	Unit       string             `bson:"unit"`
	Size       string             `bson:"size"`
	ImgPath    string             `bson:"imgPath"`
	UnitEnergy float64            `bson:"unitEnergy"`
}

type IngredientRepository struct {
	logger     *zap.SugaredLogger
	ingredColl *mongo.Collection
}

func NewIngredientRepository(logger *zap.SugaredLogger, db *mongo.Database) *IngredientRepository {
	return &IngredientRepository{logger, db.Collection(ingredCollName)}
}

func (r *IngredientRepository) Insert(ctx context.Context, ingred domain.Ingredient) (string, error) {
	res, err := r.ingredColl.InsertOne(ctx, bson.D{
		primitive.E{Key: "code", Value: ingred.Code},
		primitive.E{Key: "name", Value: ingred.Name},
		primitive.E{Key: "unit", Value: ingred.Unit},
		primitive.E{Key: "size", Value: ingred.Size},
		primitive.E{Key: "imgPath", Value: ingred.ImgPath},
		primitive.E{Key: "unitEnergy", Value: ingred.UnitEnergy},
	})
	if err != nil {
		r.logger.Errorw("failed to insert ingredient", "ingredient", ingred, "error", err)
		return "", err
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return "", fmt.Errorf("failed to convert object for ingredient %s", ingred.Name)
}

func (r *IngredientRepository) Update(ctx context.Context, id string, ingred domain.Ingredient) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Errorw("failed to convert string to mongo ObjectID", "id", id, "error", err)
		return err
	}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	res := r.ingredColl.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": bson.D{
			primitive.E{Key: "code", Value: ingred.Code},
			primitive.E{Key: "name", Value: ingred.Name},
			primitive.E{Key: "unit", Value: ingred.Unit},
			primitive.E{Key: "size", Value: ingred.Size},
			primitive.E{Key: "imgPath", Value: ingred.ImgPath},
			primitive.E{Key: "unitEnergy", Value: ingred.UnitEnergy},
		}},
		&opt,
	)

	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return &custom.ElemNotFoundError{ID: id, Err: fmt.Errorf("no ingredient found to update")}
		}
		r.logger.Errorw("failed to update ingredient", "ingredient", ingred, "error", err)
		return err
	}
	return nil
}

func (r *IngredientRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Errorw("could not convert id to objectID", "id", id, "error", err)
	}

	res, err := r.ingredColl.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		r.logger.Errorw("failed to delete ingredient", "id", id, "error", err)
		return err
	}
	if res.DeletedCount == 0 {
		return &custom.ElemNotFoundError{ID: id, Err: fmt.Errorf("no ingredient found to delete")}
	}
	return nil
}

func (r *IngredientRepository) Get(ctx context.Context, id string) (domain.Ingredient, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Errorw("could not convert id to objectID", "id", id, "error", err)
		return domain.Ingredient{}, err
	}
	var ingredDAO ingredientFetchDAO
	if err = r.ingredColl.FindOne(ctx, bson.M{"_id": objectID}).Decode(&ingredDAO); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Debugw("no ingredient found with id", "id", id)
			return domain.Ingredient{}, &custom.ElemNotFoundError{ID: id, Err: err}
		}
		r.logger.Errorw("error getting ingredient", "id", id, "error", err)
		return domain.Ingredient{}, err
	}
	return convertToDomainIngredient(ingredDAO), nil
}

func (r *IngredientRepository) GetPaginated(
	ctx context.Context, limit int, offset int, sortField string, sortDirection enum.SortDirection,
) ([]domain.Ingredient, error) {
	return nil, nil
}

func (r *IngredientRepository) GetAll(ctx context.Context) ([]domain.Ingredient, error) {
	var daos []ingredientFetchDAO
	var opts = options.Find().SetSort(bson.D{
		{Key: "code", Value: 1},
	})
	cursor, err := r.ingredColl.Find(
		ctx,
		bson.M{},
		opts,
	)
	if err != nil {
		r.logger.Errorw("failed to get ingredients from db", "error", err)
		return nil, err
	}
	if err = cursor.All(ctx, &daos); err != nil {
		r.logger.Errorw("failed to map ingredient cursor to a list of DAO objects", "error", err)
		return nil, err
	}

	ingreds := make([]domain.Ingredient, len(daos))
	for _, d := range daos {
		ingreds = append(ingreds, convertToDomainIngredient(d))
	}
	return ingreds, nil
}

// Auxiliary functions
func convertToDomainIngredient(dao ingredientFetchDAO) domain.Ingredient {
	return domain.Ingredient{
		ID:         dao.ObjectID.Hex(),
		Code:       dao.Code,
		Name:       dao.Name,
		Unit:       dao.Unit,
		Size:       dao.Size,
		ImgPath:    dao.ImgPath,
		UnitEnergy: dao.UnitEnergy,
	}
}
