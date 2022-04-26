package warehouse

import (
	"context"

	"github.com/usuario/repositorio/internal/domain"
	"github.com/usuario/repositorio/pkg/database"
)

// Repository encapsulates the storage of a warehouse.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Warehouse, error)
	Get(ctx context.Context, id int) (domain.Warehouse, error)
	Exists(ctx context.Context, warehouseCode string) bool
	Save(ctx context.Context, w domain.Warehouse) (int, error)
	Update(ctx context.Context, w domain.Warehouse) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db database.Database[domain.Warehouse]
}

func NewRepository(db database.Database[domain.Warehouse]) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	return nil, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	return domain.Warehouse{}, nil
}

func (r *repository) Exists(ctx context.Context, warehouseCode string) bool {
	return false
}

func (r *repository) Save(ctx context.Context, w domain.Warehouse) (int, error) {
	return 0, nil
}

func (r *repository) Update(ctx context.Context, w domain.Warehouse) error {
	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	return nil
}
