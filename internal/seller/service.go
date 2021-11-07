package seller

import "errors"

// Errors
var (
	ErrNotFound = errors.New("seller not found")
)

type Service interface{}

type service struct{}

func NewService() Service {
	return &service{}
}
