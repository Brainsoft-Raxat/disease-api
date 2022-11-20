package repository

import (
	"context"
	"disease-api/internal/app/config"
	"disease-api/internal/models"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type countryRepository struct {
	db  *pgxpool.Pool
	cfg config.Postgres
}

func (r *countryRepository) CreateCountry(ctx context.Context, c *models.Country) (cname string, err error) {
	q := `INSERT INTO country (cname, population) VALUES ($1, $2) RETURNING cname`

	err = r.db.QueryRow(ctx, q, c.Cname, c.Population).Scan(&cname)
	if err != nil {
		return
	}

	return
}

func (r *countryRepository) UpdateCountry(ctx context.Context, c *models.Country) error {
	q := `UPDATE country SET cname = $1, population = $2 WHERE cname = $1`

	res, err := r.db.Exec(ctx, q, c.Cname, c.Population, c.Cname)
	if err != nil {
		return err
	}

	if count := res.RowsAffected(); count == 0 {
		return fmt.Errorf("no country with cname %v", c.Cname)
	}

	return nil
}

func (r *countryRepository) DeleteCountry(ctx context.Context, cname string) error {
	q := `DELETE FROM country WHERE cname = $1`

	res, err := r.db.Exec(ctx, q, cname)
	if err != nil {
		return err
	}

	if count := res.RowsAffected(); count == 0 {
		return fmt.Errorf("no country with cname %v", cname)
	}

	return nil
}

func (r *countryRepository) GetCountry(ctx context.Context, cname string) (c *models.Country, err error) {
	q := `SELECT cname, population FROM country WHERE cname = $1`

	c = new(models.Country)
	err = r.db.QueryRow(ctx, q, cname).Scan(&c.Cname, &c.Population)
	if err != nil {
		return
	}

	return
}

func (r *countryRepository) GetAllCountries(ctx context.Context) ([]models.Country, error) {
	var res []models.Country

	q := `SELECT cname, population FROM country`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c models.Country
		err := rows.Scan(&c.Cname, &c.Population)
		if err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}

func NewCountryRepository(db *pgxpool.Pool, cfg config.Postgres) CountryRepository {
	return &countryRepository{
		db:  db,
		cfg: cfg,
	}
}
