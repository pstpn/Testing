package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
	Bucket *gridfs.Bucket
}

func New(uri, dbName, bucketName string) (*MongoDB, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	var result bson.M
	if err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	bucketOpts := options.GridFSBucket().SetName(bucketName)
	bucket, err := gridfs.NewBucket(db, bucketOpts)
	if err != nil {
		return nil, err
	}

	return &MongoDB{Client: client, Bucket: bucket}, nil
}

func (m *MongoDB) Close() {
	if m.Client != nil {
		m.Client.Disconnect(context.TODO())
	}
}
