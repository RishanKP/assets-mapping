package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Assets struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Type      string             `bson:"type" json:"type"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type Mapping struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	UserId    string             `bson:"userId" json:"userId"`
	AssetId   string             `bson:"assetId" json:"assetId"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
