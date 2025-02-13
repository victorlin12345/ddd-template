package repository

import (
	"context"
	"errors"
	"time"

	"github.com/victorlin12345/ddd-template/internal/domain/order"
	db "github.com/victorlin12345/ddd-template/internal/infrastructure/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderMongoRepo struct {
	collection *mongo.Collection
}

func NewOrderMongoRepo(mc *db.MongoClient) order.Repository {
	client := mc.Client
	collection := client.Database("test").Collection("orders")
	return &OrderMongoRepo{collection: collection}
}

// GetByID implements order.Repository.
func (o *OrderMongoRepo) GetByID(ctx context.Context, id int64) (*order.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var ord order.Order
	filter := bson.M{"_id": id}
	err := o.collection.FindOne(ctx, filter).Decode(&ord)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &ord, nil
}

func (o *OrderMongoRepo) CreateOrder(ctx context.Context, order *order.Order) (*order.Order, error) {
	res, err := o.collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	order.ID = res.InsertedID.(primitive.ObjectID)
	return order, nil
}

// Save implements order.Repository.
func (o *OrderMongoRepo) Save(ctx context.Context, ord *order.Order) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": ord.ID}
	update := bson.M{"$set": ord}
	_, err := o.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return err
}
