package repository

import (
	"asset-mapping/pkg/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MappingRepository interface {
	Create(ctx context.Context, asset models.Mapping) error
	Get(ctx context.Context, userId string) ([]models.Mapping, error)
	Delete(ctx context.Context, id string) error
	GetCountByUserId(ctx context.Context, userId string) int64
}

type mappingRepo struct {
	collection *mongo.Collection
}

func (r mappingRepo) Create(ctx context.Context, mapping models.Mapping) error {
	mapping.ID = primitive.NewObjectID()
	mapping.CreatedAt = time.Now()
	mapping.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, mapping)

	return err
}

func (r mappingRepo) Get(ctx context.Context, userId string) ([]models.Mapping, error) {

	var mappings []models.Mapping
	cur, err := r.collection.Find(ctx, bson.M{"userId": userId})
	if err != nil {
		return mappings, err
	}

	for cur.Next(context.TODO()) {
		var mapping models.Mapping

		err := cur.Decode(&mapping)
		if err != nil {
			return mappings, err
		}

		mappings = append(mappings, mapping)
	}
	return mappings, nil
}

func (r mappingRepo) Delete(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objId})

	return err
}

func (r mappingRepo) GetCountByUserId(ctx context.Context, userId string) int64 {
	count, err := r.collection.CountDocuments(ctx, bson.M{"userId": userId})
	if err != nil {
		return 0
	}

	return count
}

func NewMappingRepository(db *mongo.Database, collectionName string) MappingRepository {
	collection := db.Collection(collectionName)
	return mappingRepo{
		collection: collection,
	}
}
