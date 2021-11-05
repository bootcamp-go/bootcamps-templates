package buyer

import (
	"context"
	"database/sql"
	"errors"

	"github.com/usuario/repositorio/internal/domain"
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
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	rows, err := r.db.Query(`SELECT * FROM buyers`)
	if err != nil {
		return nil, err
	}

	var buyers []domain.Buyer

	for rows.Next() {
		b := domain.Buyer{}
		_ = rows.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName)
		buyers = append(buyers, b)
	}

	return buyers, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Buyer, error) {

	sqlStatement := `SELECT * FROM buyers WHERE id = ?;`
	row := r.db.QueryRow(sqlStatement, id)
	b := domain.Buyer{}
	err := row.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName)
	if err != nil {
		return domain.Buyer{}, err
	}

	return b, nil
}

func (r *repository) Exists(ctx context.Context, cardNumberID string) bool {
	sqlStatement := `SELECT card_number_id FROM buyers WHERE card_number_id=?;`
	row := r.db.QueryRow(sqlStatement, cardNumberID)
	err := row.Scan(&cardNumberID)
	return err == nil
}

func (r *repository) Save(ctx context.Context, b domain.Buyer) (int, error) {

	stmt, err := r.db.Prepare(`INSERT INTO buyers("card_number_id","first_name","last_name") VALUES (?,?,?)`)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&b.CardNumberID, &b.FirstName, &b.LastName)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, b domain.Buyer) error {
	stmt, err := r.db.Prepare(`UPDATE buyers SET "first_name"=?, "last_name"=?  WHERE "id"=?`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&b.FirstName, &b.LastName, &b.ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return errors.New("buyer not found")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(`DELETE FROM buyers WHERE id = ?`)
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
		return errors.New("buyer not found")
	}

	return nil
}
