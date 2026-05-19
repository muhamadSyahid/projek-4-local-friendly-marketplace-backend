package repository

import (
	"context"
	"errors"
	"pade-backend/pkg/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{
		collection: db.Collection("products"),
	}
}

func (r *ProductRepository) Create(product *entities.Product) (*entities.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}

	product.ID = result.InsertedID.(string)
	return product, nil
}

func (r *ProductRepository) GetByID(id string) (*entities.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var product entities.Product
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) GetAll() ([]*entities.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*entities.Product
	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetBySellerID(sellerID string) ([]*entities.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"sellerId": sellerID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*entities.Product
	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetByCategory(category string) ([]*entities.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"category": category})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*entities.Product
	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) Update(product *entities.Product) (*entities.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product.UpdatedAt = time.Now()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": product.ID}, bson.M{"$set": product})
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *ProductRepository) Search(query string) ([]*entities.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{
		"$or": []bson.M{
			{"name": bson.M{"$regex": query, "$options": "i"}},
			{"description": bson.M{"$regex": query, "$options": "i"}},
		},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*entities.Product
	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
