package common

import (
	"github.com/archit-p/MicroserviceTemplate/pkg/models"
)

// SampleModel is an interface for interacting with
// databases
type SampleModel interface {
	Insert(content string) (string, error)
	Get(id string) (*models.Sample, error)
	Update(id string, content string) (int64, error)
	Delete(id string) (int64, error)
	Search(keywords string) (*models.Samples, error)
	Top(parameter string, limit int64) (*models.Samples, error)
}
