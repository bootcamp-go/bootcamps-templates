package buyer

import "errors"

// Errors
var (
	ErrNotFound = errors.New("buyer not found")
)

type Service interface{}

type service struct{}

func NewService() Service {
	return &service{}
}
