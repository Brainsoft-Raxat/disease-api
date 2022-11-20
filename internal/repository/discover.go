package repository

import (
	"context"
	"disease-api/internal/app/config"
	"disease-api/internal/models"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type discoverRepository struct {
	db  *pgxpool.Pool
	cfg config.Postgres
}

func (r *discoverRepository) CreateDiscover(ctx context.Context, d *models.Discover) (code string, err error) {
	q := `INSERT INTO discover (cname, disease_code, first_enc_date) VALUES ($1, $2, $3) RETURNING disease_code`

	err = r.db.QueryRow(ctx, q, d.CName, d.DiseaseCode, d.FirstEncDate).Scan(&code)
	if err != nil {
		return
	}

	return
}

func (r *discoverRepository) UpdateDiscover(ctx context.Context, d *models.Discover) error {
	q := `UPDATE discover SET cname = $1, disease_code = $2, first_enc_date = $3 WHERE disease_code = $4`

	res, err := r.db.Exec(ctx, q, d.CName, d.DiseaseCode, d.FirstEncDate, d.DiseaseCode)
	if err != nil {
		return err
	}

	if count := res.RowsAffected(); count == 0 {
		return fmt.Errorf("no discover with disease_code %v", d.DiseaseCode)
	}

	return nil
}

func (r *discoverRepository) DeleteDiscover(ctx context.Context, country, diseaseCode string) error {
	q := `DELETE FROM discover WHERE cname = $1 AND disease_code = $2`

	res, err := r.db.Exec(ctx, q, country, diseaseCode)
	if err != nil {
		return err
	}

	if count := res.RowsAffected(); count == 0 {
		return fmt.Errorf("no discover with disease_code %v", diseaseCode)
	}

	return nil
}

func (r *discoverRepository) GetDiscover(ctx context.Context, country, diseaseCode string) (d *models.Discover, err error) {
	q := `SELECT * FROM discover WHERE cname = $1 AND disease_code = $2`

	d = new(models.Discover)
	err = r.db.QueryRow(ctx, q, country, diseaseCode).Scan(&d.CName, &d.DiseaseCode, &d.FirstEncDate)
	if err != nil {
		return
	}

	return
}

func (r *discoverRepository) GetAllDiscovers(ctx context.Context) ([]models.Discover, error) {
	var res []models.Discover

	q := `SELECT * FROM discover`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var d models.Discover
		err := rows.Scan(&d.CName, &d.DiseaseCode, &d.FirstEncDate)
		if err != nil {
			return nil, err
		}

		res = append(res, d)
	}

	return res, nil
}

func NewDiscoverRepository(db *pgxpool.Pool, cfg config.Postgres) DiscoverRepository {
	return &discoverRepository{
		db:  db,
		cfg: cfg,
	}
}
