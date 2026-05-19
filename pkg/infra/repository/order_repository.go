package repository

import (
	"context"
	"errors"
	"pade-backend/pkg/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{
		collection: db.Collection("orders"),
	}
}

func (r *OrderRepository) Create(order *entities.Order) (*entities.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}

	order.ID = result.InsertedID.(string)
	return order, nil
}

func (r *OrderRepository) GetByID(id string) (*entities.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var order entities.Order
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) GetByBuyerID(buyerID string) ([]*entities.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"userId": buyerID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []*entities.Order
	err = cursor.All(ctx, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) GetBySellerID(sellerID string) ([]*entities.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"sellerId": sellerID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []*entities.Order
	err = cursor.All(ctx, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) Update(order *entities.Order) (*entities.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	order.UpdatedAt = time.Now()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": order.ID}, bson.M{"$set": order})
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *OrderRepository) GetAll() ([]*entities.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []*entities.Order
	err = cursor.All(ctx, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
