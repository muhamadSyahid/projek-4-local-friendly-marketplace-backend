package repository

import (
	"context"
	"errors"
	"pade-backend/pkg/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MarketplaceRepository struct {
	collection *mongo.Collection
}

func NewMarketplaceRepository(db *mongo.Database) *MarketplaceRepository {
	return &MarketplaceRepository{
		collection: db.Collection("marketplaces"),
	}
}

func (r *MarketplaceRepository) Create(marketplace *entities.Marketplace) (*entities.Marketplace, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	marketplace.CreatedAt = time.Now()
	marketplace.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, marketplace)
	if err != nil {
		return nil, err
	}

	marketplace.ID = result.InsertedID.(string)
	return marketplace, nil
}

func (r *MarketplaceRepository) GetByID(id string) (*entities.Marketplace, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var marketplace entities.Marketplace
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&marketplace)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("marketplace not found")
		}
		return nil, err
	}

	return &marketplace, nil
}

func (r *MarketplaceRepository) GetAll() ([]*entities.Marketplace, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var marketplaces []*entities.Marketplace
	err = cursor.All(ctx, &marketplaces)
	if err != nil {
		return nil, err
	}

	return marketplaces, nil
}

func (r *MarketplaceRepository) GetByOwnerID(ownerID string) ([]*entities.Marketplace, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"ownerId": ownerID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var marketplaces []*entities.Marketplace
	err = cursor.All(ctx, &marketplaces)
	if err != nil {
		return nil, err
	}

	return marketplaces, nil
}

func (r *MarketplaceRepository) GetByCategory(category string) ([]*entities.Marketplace, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"category": category})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var marketplaces []*entities.Marketplace
	err = cursor.All(ctx, &marketplaces)
	if err != nil {
		return nil, err
	}

	return marketplaces, nil
}

func (r *MarketplaceRepository) Update(marketplace *entities.Marketplace) (*entities.Marketplace, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	marketplace.UpdatedAt = time.Now()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": marketplace.ID}, bson.M{"$set": marketplace})
	if err != nil {
		return nil, err
	}

	return marketplace, nil
}

func (r *MarketplaceRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
