package repository

import (
	"asset-mapping/pkg/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) error
	Update(ctx context.Context, user models.User) error
	GetByEmail(ctx context.Context, email string) (models.User, error)
	GetById(ctx context.Context, id string) (models.User, error)
	Get(ctx context.Context) ([]models.User, error)
	Delete(ctx context.Context, id string) error
}

type userRepo struct {
	collection *mongo.Collection
}

func (r userRepo) Create(ctx context.Context, user models.User) error {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, user)

	return err
}

func (r userRepo) Update(ctx context.Context, user models.User) error {
	user.UpdatedAt = time.Now()
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err
}

func (r userRepo) GetByEmail(ctx context.Context, email string) (models.User, error) {

	filter := bson.M{"email": email}

	var user models.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r userRepo) GetById(ctx context.Context, id string) (models.User, error) {

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}
	filter := bson.M{"_id": objId}

	var user models.User
	err = r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r userRepo) Get(ctx context.Context) ([]models.User, error) {

	var users []models.User
	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return users, err
	}

	for cur.Next(context.TODO()) {
		var user models.User

		err := cur.Decode(&user)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (r userRepo) Delete(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objId})

	return err
}

func NewUserRepository(db *mongo.Database, collectionName string) UserRepository {
	collection := db.Collection(collectionName)
	return userRepo{
		collection: collection,
	}
}
