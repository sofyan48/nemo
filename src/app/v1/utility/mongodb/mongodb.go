package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB ...
type MongoDB struct{}

// MongoDBHandler ...
func MongoDBHandler() *MongoDB {
	return &MongoDB{}
}

// MongoDBInterface ...
type MongoDBInterface interface {
	InsertOne(collection string, data interface{}) (*mongo.InsertOneResult, error)
	InsertMany(collection string, data []interface{}) (*mongo.InsertManyResult, error)
	Find(collection string, filter interface{}) (*mongo.Cursor, context.Context, error)
	Delete(collection string, filter interface{}) (*mongo.DeleteResult, error)
	Edit(collection string, filter, data interface{}) (*mongo.UpdateResult, error)
}

func (mongolib *MongoDB) init() (*mongo.Database, context.Context, error) {
	mongoAddres := os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT")
	ctx := context.Background()
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://" + mongoAddres)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, nil, err
	}
	return client.Database(os.Getenv("MONGO_DATABASE")), ctx, nil
}

// InsertOne ...
func (mongolib *MongoDB) InsertOne(collection string, data interface{}) (*mongo.InsertOneResult, error) {
	db, ctx, err := mongolib.init()
	if err != nil {
		return nil, err
	}

	return db.Collection(collection).InsertOne(ctx, data)
}

// InsertMany ...
func (mongolib *MongoDB) InsertMany(collection string, data []interface{}) (*mongo.InsertManyResult, error) {
	db, ctx, err := mongolib.init()
	if err != nil {
		return nil, err
	}
	return db.Collection(collection).InsertMany(ctx, data)
}

// Find ...
func (mongolib *MongoDB) Find(collection string, filter interface{}) (*mongo.Cursor, context.Context, error) {
	db, ctx, err := mongolib.init()
	if err != nil {
		return nil, nil, err
	}
	csr, err := db.Collection(collection).Find(ctx, filter)
	if err != nil {
		defer csr.Close(ctx)
		return nil, nil, err
	}
	return csr, ctx, nil
}

// Delete ..
func (mongolib *MongoDB) Delete(collection string, filter interface{}) (*mongo.DeleteResult, error) {
	db, ctx, err := mongolib.init()
	if err != nil {
		return nil, err
	}
	return db.Collection(collection).DeleteOne(ctx, filter)
}

// Edit ...
func (mongolib *MongoDB) Edit(collection string, filter, data interface{}) (*mongo.UpdateResult, error) {
	db, ctx, err := mongolib.init()
	if err != nil {
		return nil, err
	}
	return db.Collection(collection).UpdateOne(ctx, filter, bson.M{"$set": data})
}
