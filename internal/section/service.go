package section

import "errors"

// Errors
var (
	ErrNotFound = errors.New("section not found")
)

type Service interface{}

type service struct{}

func NewService() Service {
	return &service{}
}
