package mongodb

import (
	"context"
	"time"

	"github.com/archit-p/MicroserviceTemplate/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SampleMongo is a an implementation for the SampleModel
// interface to interact with Samples in MongoDB
type SampleMongo struct {
	Collection *mongo.Collection
}

// NewSampleMongo initializes the Mongo connection
func NewSampleMongo(client *mongo.Client) (*SampleMongo, error) {
	s := SampleMongo{
		Collection: client.Database("example").Collection("sample"),
	}
	index := mongo.IndexModel{
		Keys: bson.D{
			bson.E{Key: "content", Value: "text"},
		},
	}

	_, err := s.Collection.Indexes().CreateOne(context.TODO(), index)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// Insert (SampleMongo) inserts a new Sample into the database
// returns the id of new entry and an error
func (s *SampleMongo) Insert(content string) (string, error) {
	snip := &models.Sample{
		Content: content,
		Likes:   0,
		Created: time.Now(),
		Deleted: false,
	}

	res, err := s.Collection.InsertOne(context.TODO(), snip)
	if err != nil {
		return "", err
	}
	ID := res.InsertedID.(primitive.ObjectID).Hex()

	return ID, nil
}

// Get (SampleMongo) gets a Sample based on the ID
func (s *SampleMongo) Get(id string) (*models.Sample, error) {
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{bson.E{Key: "_id", Value: mongoID}, bson.E{Key: "deleted", Value: false}}

	var res models.Sample
	err = s.Collection.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Update (SampleMongo) updates a Sample based on its ID
// To-do
func (s *SampleMongo) Update(id string, content string) (int64, error) {
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	filter := bson.D{bson.E{Key: "_id", Value: mongoID}, bson.E{Key: "deleted", Value: false}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "content", Value: content}}}}

	res, err := s.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0, err
	}

	return res.ModifiedCount, nil
}

// Delete (SampleMongo) deletes a Sample based on its ID
// To-do
func (s *SampleMongo) Delete(id string) (int64, error) {
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	filter := bson.D{bson.E{Key: "_id", Value: mongoID}, bson.E{Key: "deleted", Value: false}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "deleted", Value: true}}}}

	res, err := s.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0, err
	}

	return res.ModifiedCount, nil
}

// Search (SampleMongo) returns Samples based on keywords
func (s *SampleMongo) Search(keywords string) (*models.Samples, error) {
	filter := bson.M{"$text": bson.M{"$search": keywords}}
	opts := options.Find().SetSort(bson.M{"score": bson.M{"$meta": "textScore"}})

	cur, err := s.Collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var res models.Samples
	for cur.Next(context.TODO()) {
		var resi models.Sample
		err = cur.Decode(&resi)
		if err != nil {
			return nil, err
		}
		res = append(res, resi)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &res, nil
}

// Top (SampleMongo) returns the top Samples based on search field and limit
func (s *SampleMongo) Top(parameter string, limit int64) (*models.Samples, error) {
	opts := options.Find().SetSort(bson.M{parameter: 1}).SetLimit(limit)
	filter := bson.D{}

	cur, err := s.Collection.Find(context.TODO(), filter, opts)

	if err != nil {
		return nil, err
	}

	var res models.Samples
	for cur.Next(context.TODO()) {
		var resi models.Sample
		err = cur.Decode(&resi)
		if err != nil {
			return nil, err
		}
		res = append(res, resi)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &res, nil
}
