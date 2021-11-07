package employee

import "errors"

// Errors
var (
	ErrNotFound = errors.New("employee not found")
)

type Service interface{}

type service struct{}

func NewService() Service {
	return &service{}
}
