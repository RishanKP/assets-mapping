package repository

import (
	"asset-mapping/pkg/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AssetsRepository interface {
	Create(ctx context.Context, asset models.Assets) error
	Update(ctx context.Context, asset models.Assets) error
	GetById(ctx context.Context, id string) (models.Assets, error)
	Get(ctx context.Context) ([]models.Assets, error)
	Delete(ctx context.Context, id string) error
}

type assetsRepo struct {
	collection *mongo.Collection
}

func (r assetsRepo) Create(ctx context.Context, asset models.Assets) error {
	asset.ID = primitive.NewObjectID()
	asset.CreatedAt = time.Now()
	asset.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, asset)

	return err
}

func (r assetsRepo) Update(ctx context.Context, asset models.Assets) error {
	asset.UpdatedAt = time.Now()
	filter := bson.M{"_id": asset.ID}
	update := bson.M{"$set": asset}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err
}

func (r assetsRepo) GetById(ctx context.Context, id string) (models.Assets, error) {

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Assets{}, err
	}
	filter := bson.M{"_id": objId}

	var asset models.Assets
	err = r.collection.FindOne(ctx, filter).Decode(&asset)
	if err != nil {
		return asset, err
	}

	return asset, nil
}

func (r assetsRepo) Get(ctx context.Context) ([]models.Assets, error) {

	var assets []models.Assets
	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return assets, err
	}

	for cur.Next(context.TODO()) {
		var asset models.Assets

		err := cur.Decode(&asset)
		if err != nil {
			return assets, err
		}

		assets = append(assets, asset)
	}
	return assets, nil
}

func (r assetsRepo) Delete(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objId})

	return err
}

func NewAssetsRepository(db *mongo.Database, collectionName string) AssetsRepository {
	collection := db.Collection(collectionName)
	return assetsRepo{
		collection: collection,
	}
}
