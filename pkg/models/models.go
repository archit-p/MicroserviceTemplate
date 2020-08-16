package models

import (
	"time"
)

// Sample represents a simple piece of data
type Sample struct {
	Content string    `bson:"content"`
	Likes   int       `bson:"likes"`
	Created time.Time `bson:"created"`
	Deleted bool      `bson:"deleted"`
}

// Samples are a slice of Sample
type Samples []Sample

// SampleModel is an interface for interacting with
// databases
type SampleModel interface {
	Insert(content string) (string, error)
	Get(id string) (*Sample, error)
	Update(id string, content string) (int64, error)
	Delete(id string) (int64, error)
	Search(keywords string) (*Samples, error)
	Top(parameter string, limit int64) (*Samples, error)
}
