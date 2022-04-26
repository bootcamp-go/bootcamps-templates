package section

import (
	"context"

	"github.com/usuario/repositorio/internal/domain"
	"github.com/usuario/repositorio/pkg/database"
)

// Repository encapsulates the storage of a section.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Section, error)
	Get(ctx context.Context, id int) (domain.Section, error)
	Exists(ctx context.Context, cid int) bool
	Save(ctx context.Context, s domain.Section) (int, error)
	Update(ctx context.Context, s domain.Section) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db database.Database[domain.Section]
}

func NewRepository(db database.Database[domain.Section]) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Section, error) {
	return nil, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Section, error) {
	return domain.Section{}, nil
}

func (r *repository) Exists(ctx context.Context, sectionNumber int) bool {
	return false
}

func (r *repository) Save(ctx context.Context, s domain.Section) (int, error) {
	return 0, nil
}

func (r *repository) Update(ctx context.Context, s domain.Section) error {
	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	return nil
}
