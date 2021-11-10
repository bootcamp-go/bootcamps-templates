package seller

import (
	"context"
	"database/sql"

	"github.com/usuario/repositorio/internal/domain"
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
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Seller, error) {
	query := "SELECT * FROM sellers"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var sellers []domain.Seller

	for rows.Next() {
		s := domain.Seller{}
		_ = rows.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityID)
		sellers = append(sellers, s)
	}

	return sellers, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Seller, error) {
	query := "SELECT * FROM sellers WHERE id=?;"
	row := r.db.QueryRow(query, id)
	s := domain.Seller{}
	err := row.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityID)
	if err != nil {
		return domain.Seller{}, err
	}

	return s, nil
}

func (r *repository) Exists(ctx context.Context, cid int) bool {
	query := "SELECT cid FROM sellers WHERE cid=?;"
	row := r.db.QueryRow(query, cid)
	err := row.Scan(&cid)
	return err == nil
}

func (r *repository) Save(ctx context.Context, s domain.Seller) (int, error) {
	query := "INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, s domain.Seller) error {
	query := "UPDATE sellers SET cid=?, company_name=?, address=?, telephone=?, locality_id=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID, s.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM sellers WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return ErrNotFound
	}

	return nil
}
