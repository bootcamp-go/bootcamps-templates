package product

import "errors"

// Errors
var (
	ErrNotFound = errors.New("product not found")
)

type Service interface{}

type service struct{}

func NewService() Service {
	return &service{}
}
