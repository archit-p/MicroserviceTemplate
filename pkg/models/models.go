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
