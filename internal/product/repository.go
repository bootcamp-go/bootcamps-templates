package product

import (
	"context"

	"github.com/usuario/repositorio/internal/domain"
	"github.com/usuario/repositorio/pkg/database"
)

// Repository encapsulates the storage of a Product.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Exists(ctx context.Context, productCode string) bool
	Save(ctx context.Context, p domain.Product) (int, error)
	Update(ctx context.Context, p domain.Product) error
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

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	return nil, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Product, error) {
	return domain.Product{}, nil
}

func (r *repository) Exists(ctx context.Context, productCode string) bool {
	return false
}

func (r *repository) Save(ctx context.Context, p domain.Product) (int, error) {
	return 0, nil
}

func (r *repository) Update(ctx context.Context, p domain.Product) error {
	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	return nil
}
