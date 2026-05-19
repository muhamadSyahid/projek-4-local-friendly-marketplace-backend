package repository

import (
	"context"
	"errors"
	"pade-backend/pkg/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SellerRepository struct {
	collection *mongo.Collection
}

func NewSellerRepository(db *mongo.Database) *SellerRepository {
	return &SellerRepository{
		collection: db.Collection("sellers"),
	}
}

func (r *SellerRepository) Create(seller *entities.Seller) (*entities.Seller, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	seller.CreatedAt = time.Now()
	seller.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, seller)
	if err != nil {
		return nil, err
	}

	seller.ID = result.InsertedID.(string)
	return seller, nil
}

func (r *SellerRepository) GetByID(id string) (*entities.Seller, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var seller entities.Seller
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&seller)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("seller not found")
		}
		return nil, err
	}

	return &seller, nil
}

func (r *SellerRepository) GetByUserID(userID string) (*entities.Seller, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var seller entities.Seller
	err := r.collection.FindOne(ctx, bson.M{"userId": userID}).Decode(&seller)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("seller not found")
		}
		return nil, err
	}

	return &seller, nil
}

func (r *SellerRepository) GetAll() ([]*entities.Seller, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var sellers []*entities.Seller
	err = cursor.All(ctx, &sellers)
	if err != nil {
		return nil, err
	}

	return sellers, nil
}

func (r *SellerRepository) Update(seller *entities.Seller) (*entities.Seller, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	seller.UpdatedAt = time.Now()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": seller.ID}, bson.M{"$set": seller})
	if err != nil {
		return nil, err
	}

	return seller, nil
}

func (r *SellerRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *SellerRepository) GetNearestStores(latitude, longitude float64, limit int) ([]*entities.Seller, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if limit <= 0 {
		limit = 10
	}

	opts := options.Find().SetLimit(int64(limit))

	cursor, err := r.collection.Find(ctx, bson.M{
		"storeLocation": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{longitude, latitude},
				},
			},
		},
	}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var sellers []*entities.Seller
	err = cursor.All(ctx, &sellers)
	if err != nil {
		return nil, err
	}

	return sellers, nil
}
