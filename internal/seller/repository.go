package seller

import (
	"context"

	"github.com/usuario/repositorio/internal/domain"
	"github.com/usuario/repositorio/pkg/database"
)

// Repository encapsulates the storage of a Seller.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Get(ctx context.Context, id int) (domain.Seller, error)
	Exists(ctx context.Context, cid int) bool
	Save(ctx context.Context, s domain.Seller) (int, error)
	Update(ctx context.Context, s domain.Seller) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db database.Database
}

func NewRepository(db database.Database) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Seller, error) {
	return nil, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Seller, error) {
	return domain.Seller{}, nil
}

func (r *repository) Exists(ctx context.Context, cid int) bool {
	return false
}

func (r *repository) Save(ctx context.Context, s domain.Seller) (int, error) {
	return 0, nil
}

func (r *repository) Update(ctx context.Context, s domain.Seller) error {
	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	return nil
}
