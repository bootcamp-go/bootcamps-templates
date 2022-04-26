package buyer

import (
	"context"

	"github.com/usuario/repositorio/internal/domain"
	"github.com/usuario/repositorio/pkg/database"
)

// Repository encapsulates the storage of a buyer.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Buyer, error)
	Get(ctx context.Context, id int) (domain.Buyer, error)
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, b domain.Buyer) (int, error)
	Update(ctx context.Context, b domain.Buyer) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db database.Database[domain.Buyer]
}

func NewRepository(db database.Database[domain.Buyer]) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	return nil, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Buyer, error) {
	return domain.Buyer{}, nil
}

func (r *repository) Exists(ctx context.Context, cardNumberID string) bool {
	return false
}

func (r *repository) Save(ctx context.Context, b domain.Buyer) (int, error) {
	return 0, nil
}

func (r *repository) Update(ctx context.Context, b domain.Buyer) error {
	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	return nil
}
