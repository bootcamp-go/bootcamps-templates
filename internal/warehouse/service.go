package warehouse

import "errors"

// Errors
var (
	ErrNotFound = errors.New("warehouse not found")
)

type Service interface{}

type service struct{}

func NewService() Service {
	return &service{}
}
