package db

import (
	"context"
	"fmt"
	"time"

	"github.com/gimmickless/keat-kit-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const (
	kitCollName = "kits"
)

type kitPriceDAO struct {
	Amount   float64 `bson:"amount"`
	Country  string  `bson:"country"`
	Currency string  `bson:"currency"`
}

type kitPreparationStepDAO struct {
	IngredientId primitive.ObjectID `bson:"ingredientId"`
	Quantity     float64            `bson:"quantity"`
	Action       string             `bson:"action"`
}

type kitFetchDAO struct {
	ObjectID      primitive.ObjectID      `bson:"_id"`
	CategoryIDs   []string                `bson:"categoryIds"`
	Name          string                  `bson:"name"`
	Version       string                  `bson:"version"`
	Description   string                  `bson:"description"`
	Status        string                  `bson:"status"`
	AuthorID      string                  `bson:"authorId"`
	Recipe        []kitPreparationStepDAO `bson:"recipe"`
	Energy        float64                 `bson:"energy"`
	Portion       float64                 `bson:"portion"`
	PrepTime      int                     `bson:"prepTime"`
	LastUpdatedAt time.Time               `bson:"lastUpdatedAt"`
	Prices        []kitPriceDAO           `bson:"prices"`
}

type KitRepository struct {
	logger  *zap.SugaredLogger
	kitColl *mongo.Collection
}

func NewKitRepository(logger *zap.SugaredLogger, db *mongo.Database) *KitRepository {
	return &KitRepository{logger, db.Collection(kitCollName)}
}

func (r *KitRepository) Insert(ctx context.Context, kit domain.Kit) (string, error) {
	recipeDAO, err := convertToDAORecipe(kit.Recipe)
	if err != nil {
		r.logger.Errorw("failed to convert to kit DAO", "kit.recipe", kit.Recipe, "error", err)
		return "", err
	}
	pricesDAO, err := convertToDAOPrices(kit.Prices)
	if err != nil {
		r.logger.Errorw("failed to convert to kit DAO", "kit.prices", kit.Prices, "error", err)
		return "", err
	}

	res, err := r.kitColl.InsertOne(ctx, bson.D{
		primitive.E{Key: "categoryIds", Value: kit.CategoryIDs},
		primitive.E{Key: "name", Value: kit.Name},
		primitive.E{Key: "version", Value: kit.Version},
		primitive.E{Key: "description", Value: kit.Description},
		primitive.E{Key: "status", Value: kit.Status},
		primitive.E{Key: "authorId", Value: kit.AuthorID},
		primitive.E{Key: "recipe", Value: recipeDAO},
		primitive.E{Key: "energy", Value: kit.Energy},
		primitive.E{Key: "portion", Value: kit.Portion},
		primitive.E{Key: "prepTime", Value: kit.PrepTime},
		primitive.E{Key: "lastUpdatedAt", Value: kit.LastUpdatedAt},
		primitive.E{Key: "prices", Value: pricesDAO},
	})
	if err != nil {
		r.logger.Errorw("failed to insert kit", "kit", kit, "error", err)
		return "", err
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return "", fmt.Errorf("failed to convert object for kit %s", kit.Name)
}

// Updates the non-price fields of the kit
func (r *KitRepository) Update(ctx context.Context, id string, kit domain.Kit) error {
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	recipeDAO, err := convertToDAORecipe(kit.Recipe)
	if err != nil {
		r.logger.Errorw("failed to convert to kit DAO", "kit.recipe", kit.Recipe, "error", err)
		return err
	}

	res := r.kitColl.FindOneAndUpdate(
		ctx,
		bson.M{"_id": id},
		bson.D{
			primitive.E{Key: "categoryIds", Value: kit.CategoryIDs},
			primitive.E{Key: "name", Value: kit.Name},
			primitive.E{Key: "version", Value: kit.Version},
			primitive.E{Key: "description", Value: kit.Description},
			primitive.E{Key: "status", Value: kit.Status},
			primitive.E{Key: "authorId", Value: kit.AuthorID},
			primitive.E{Key: "recipe", Value: recipeDAO},
			primitive.E{Key: "energy", Value: kit.Energy},
			primitive.E{Key: "portion", Value: kit.Portion},
			primitive.E{Key: "prepTime", Value: kit.PrepTime},
			primitive.E{Key: "lastUpdatedAt", Value: kit.LastUpdatedAt},
		},
		&opt,
	)

	if err := res.Err(); err != nil {
		r.logger.Errorw("failed to update kit", "kit", kit, "error", err)
		return err
	}
	return nil
}

// Updates the prices of the kit
func (r *KitRepository) UpdatePrice(ctx context.Context, id string, prices []domain.Price) error {
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	priceDAOs, err := convertToDAOPrices(prices)
	if err != nil {
		r.logger.Errorw("failed to convert to kit DAO prices", "prices", prices, "error", err)
		return err
	}

	res := r.kitColl.FindOneAndUpdate(
		ctx,
		bson.M{"_id": id},
		bson.D{
			primitive.E{Key: "prices", Value: priceDAOs},
		},
		&opt,
	)

	if err := res.Err(); err != nil {
		r.logger.Errorw("failed to update kit prices", "prices", prices, "error", err)
		return err
	}
	return nil
}

func (r *KitRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Errorw("could not convert id to objectID", "id", id, "error", err)
	}

	res, err := r.kitColl.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		r.logger.Errorw("failed to delete kit", "id", id, "error", err)
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no kit found to delete with id %s", id)
	}
	return nil
}

func (r *KitRepository) Get(ctx context.Context, id string) (domain.Kit, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Errorw("could not convert id to objectID", "id", id, "error", err)
		return domain.Kit{}, err
	}
	var kitDAO kitFetchDAO
	if err = r.kitColl.FindOne(ctx, bson.M{"_id": objectID}).Decode(&kitDAO); err != nil {
		r.logger.Errorw("error getting kit", "id", id, "error", err)
		return domain.Kit{}, err
	}
	return convertToDomainKit(kitDAO), nil
}

func (r *KitRepository) GetAll(ctx context.Context) ([]domain.Kit, error) {
	var daos []kitFetchDAO
	var opts = options.Find().SetSort(bson.D{
		{Key: "name", Value: 1},
	})
	cursor, err := r.kitColl.Find(
		ctx,
		bson.M{},
		opts,
	)
	if err != nil {
		r.logger.Errorw("failed to get kits from db", "error", err)
		return nil, err
	}
	if err = cursor.All(ctx, &daos); err != nil {
		r.logger.Errorw("failed to map kit cursor to a list of DAO objects", "error", err)
		return nil, err
	}

	kits := make([]domain.Kit, len(daos))
	for _, d := range daos {
		kits = append(kits, convertToDomainKit(d))
	}
	return kits, nil
}

// Auxiliary functions
func convertToDomainKit(dao kitFetchDAO) domain.Kit {
	recipe := convertToDomainRecipe(dao.Recipe)
	prices := convertToDomainPrices(dao.Prices)
	return domain.Kit{
		ID:            dao.ObjectID.Hex(),
		CategoryIDs:   dao.CategoryIDs,
		Name:          dao.Name,
		Version:       dao.Version,
		Description:   dao.Description,
		Status:        dao.Status,
		AuthorID:      dao.AuthorID,
		Recipe:        recipe,
		Energy:        dao.Energy,
		Portion:       dao.Portion,
		PrepTime:      dao.PrepTime,
		CreatedAt:     dao.ObjectID.Timestamp(),
		LastUpdatedAt: dao.LastUpdatedAt,
		Prices:        prices,
	}
}

func convertToDAORecipe(recipe []domain.PreparationStep) ([]kitPreparationStepDAO, error) {
	res := make([]kitPreparationStepDAO, len(recipe))
	for _, s := range recipe {
		ingredID, err := primitive.ObjectIDFromHex(s.Ingredient.ID)
		if err != nil {
			return nil, err
		}
		res = append(res, kitPreparationStepDAO{
			IngredientId: ingredID,
			Quantity:     s.Quantity,
			Action:       s.Action,
		})
	}
	return res, nil
}

func convertToDomainRecipe(recipeDAO []kitPreparationStepDAO) []domain.PreparationStep {
	res := make([]domain.PreparationStep, len(recipeDAO))
	for _, s := range recipeDAO {
		res = append(res, domain.PreparationStep{
			Ingredient: domain.Ingredient{
				ID: s.IngredientId.Hex(),
			},
			Quantity: s.Quantity,
			Action:   s.Action,
		})
	}
	return res
}

func convertToDAOPrices(prices []domain.Price) ([]kitPriceDAO, error) {
	res := make([]kitPriceDAO, len(prices))
	for _, p := range prices {
		res = append(res, kitPriceDAO{
			Amount:   p.Amount,
			Country:  p.Country,
			Currency: p.Currency,
		})
	}
	return res, nil
}

func convertToDomainPrices(pricesDAO []kitPriceDAO) []domain.Price {
	res := make([]domain.Price, len(pricesDAO))
	for _, p := range pricesDAO {
		res = append(res, domain.Price{
			Amount:   p.Amount,
			Country:  p.Country,
			Currency: p.Currency,
		})
	}
	return res
}
