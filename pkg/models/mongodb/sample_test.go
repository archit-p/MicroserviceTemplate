package mongodb

import (
	"context"
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestIT_Insert(t *testing.T) {
	dbURL := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(dbURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		t.Fatal(err)
	}

	s, err := NewSampleMongo(client)
	if err != nil {
		t.Fatal(err)
	}

	testContent := "hello, testing!"
	_, err = s.Insert(testContent)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIT_Get(t *testing.T) {
	dbURL := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(dbURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		t.Fatal(err)
	}

	s, err := NewSampleMongo(client)
	if err != nil {
		t.Fatal(err)
	}

	testContent := "hello, testing!"
	id, err := s.Insert(testContent)
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.Get(id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIT_Get_NonExistentID(t *testing.T) {
	dbURL := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(dbURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		t.Fatal(err)
	}

	s, err := NewSampleMongo(client)
	if err != nil {
		t.Fatal(err)
	}

	id := "1234"
	_, err = s.Get(id)
	if err == nil {
		t.Fatal(errors.New("ID doesn't exist"))
	}
}

func TestIT_Update(t *testing.T) {
	dbURL := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(dbURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		t.Fatal(err)
	}

	s, err := NewSampleMongo(client)
	if err != nil {
		t.Fatal(err)
	}

	testContent := "hello, testing!"
	id, err := s.Insert(testContent)
	if err != nil {
		t.Fatal(err)
	}

	updateContent := "hello, testing updated!"
	res, err := s.Update(id, updateContent)
	if err != nil || res == 0 {
		t.Fatal(err)
	}
}

func TestIT_Update_NonExistentID(t *testing.T) {
	dbURL := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(dbURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		t.Fatal(err)
	}

	s, err := NewSampleMongo(client)
	if err != nil {
		t.Fatal(err)
	}

	id := "1234"
	updateContent := "hello, testing updated!"
	res, err := s.Update(id, updateContent)
	if err == nil || res == 1 {
		t.Fatal(errors.New("ID doesn't exist"))
	}
}

func TestIT_Delete(t *testing.T) {
	dbURL := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(dbURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		t.Fatal(err)
	}

	s, err := NewSampleMongo(client)
	if err != nil {
		t.Fatal(err)
	}

	testContent := "hello, testing!"
	id, err := s.Insert(testContent)
	if err != nil {
		t.Fatal(err)
	}

	res, err := s.Delete(id)
	if err != nil || res == 0 {
		t.Fatal(err)
	}
}

func TestIT_Delete_NonExistentID(t *testing.T) {
	dbURL := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(dbURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		t.Fatal(err)
	}

	s, err := NewSampleMongo(client)
	if err != nil {
		t.Fatal(err)
	}

	id := "1234"
	res, err := s.Delete(id)
	if err == nil || res != 0 {
		t.Fatal(errors.New("ID doesn't exist"))
	}
}

func TestIT_Search(t *testing.T) {
	dbURL := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(dbURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		t.Fatal(err)
	}

	s, err := NewSampleMongo(client)
	if err != nil {
		t.Fatal(err)
	}

	keywords := "hello"
	_, err = s.Search(keywords)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIT_Top(t *testing.T) {
	dbURL := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(dbURL)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		t.Fatal(err)
	}

	s, err := NewSampleMongo(client)
	if err != nil {
		t.Fatal(err)
	}

	param := "likes"
	limit := int64(10)
	_, err = s.Top(param, limit)
	if err != nil {
		t.Fatal(err)
	}
}
