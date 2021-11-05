package warehouse

import (
	"context"
	"database/sql"
	"errors"

	"github.com/usuario/repositorio/internal/domain"
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
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	rows, err := r.db.Query(`SELECT * FROM warehouses`)
	if err != nil {
		return nil, err
	}

	var warehouses []domain.Warehouse

	for rows.Next() {
		w := domain.Warehouse{}
		_ = rows.Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimunCapacity, &w.MinimunTemperature, &w.SectionNumber)
		warehouses = append(warehouses, w)
	}

	return warehouses, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Warehouse, error) {

	sqlStatement := `SELECT * FROM warehouses WHERE id=?;`
	row := r.db.QueryRow(sqlStatement, id)
	w := domain.Warehouse{}
	err := row.Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimunCapacity, &w.MinimunTemperature, &w.SectionNumber)
	if err != nil {
		return domain.Warehouse{}, err
	}

	return w, nil
}

func (r *repository) Exists(ctx context.Context, warehouseCode string) bool {
	sqlStatement := `SELECT warehouse_code FROM warehouses WHERE "warehouse_code"=?;`
	row := r.db.QueryRow(sqlStatement, warehouseCode)
	err := row.Scan(&warehouseCode)
	return err == nil
}

func (r *repository) Save(ctx context.Context, w domain.Warehouse) (int, error) {

	stmt, err := r.db.Prepare(`INSERT INTO warehouses("address","telephone","warehouse_code","minimun_capacity","minimun_temperature", "section_number") VALUES (?,?,?,?,?,?)`)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimunCapacity, &w.MinimunTemperature, &w.SectionNumber)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, w domain.Warehouse) error {
	stmt, err := r.db.Prepare(`UPDATE warehouses SET "address"=?, "telephone"=?, "warehouse_code"=?, "minimun_capacity"=?, "minimun_temperature"=?, "section_number"=?  WHERE "id"=?`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimunCapacity, &w.MinimunTemperature, &w.SectionNumber, &w.ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return errors.New("warehouse not found")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(`DELETE FROM warehouses WHERE id=?`)
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
		return errors.New("warehouse not found")
	}

	return nil
}
