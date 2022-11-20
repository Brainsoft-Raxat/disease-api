package repository

import (
	"context"
	"disease-api/internal/app/config"
	"disease-api/internal/models"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type diseaseRepository struct {
	db  *pgxpool.Pool
	cfg config.Postgres
}

func (r *diseaseRepository) CreateDisease(ctx context.Context, d *models.Disease) (code string, err error) {
	q := `INSERT INTO disease (code, pathogen, description, id) VALUES ($1, $2, $3, $4) RETURNING code`

	err = r.db.QueryRow(ctx, q, d.Description, d.Pathogen, d.Description, d.DiseaseTypeID).Scan(&code)
	if err != nil {
		return
	}

	return
}

func (r *diseaseRepository) UpdateDisease(ctx context.Context, d *models.Disease) error {
	q := `UPDATE disease SET code = $1, pathogen = $2, description = $3, id = $4 WHERE code = $5`

	res, err := r.db.Exec(ctx, q, d.Code, d.Pathogen, d.Description, d.DiseaseTypeID, d.Code)
	if err != nil {
		return err
	}

	if count := res.RowsAffected(); count == 0 {
		return fmt.Errorf("no disease with code %v", d.Code)
	}

	return nil
}

func (r *diseaseRepository) DeleteDisease(ctx context.Context, code string) error {
	q := `DELETE FROM disease WHERE code = $1`

	res, err := r.db.Exec(ctx, q, code)
	if err != nil {
		return err
	}

	if count := res.RowsAffected(); count == 0 {
		return fmt.Errorf("no disease with code %v", code)
	}

	return nil
}

func (r *diseaseRepository) GetDisease(ctx context.Context, code string) (d *models.Disease, err error) {
	q := `SELECT d.code, d.pathogen, d.description, d.id, dt.description FROM disease d
                                       INNER JOIN disease_type dt ON d.id = dt.id 
                                       WHERE code = $1`

	d = new(models.Disease)
	err = r.db.QueryRow(ctx, q, code).Scan(&d.Code, &d.Pathogen, &d.Description, &d.DiseaseTypeID, &d.DiseaseType)
	if err != nil {
		return
	}

	return
}

func (r *diseaseRepository) GetAllDiseases(ctx context.Context) ([]models.Disease, error) {
	var res []models.Disease

	q := `SELECT d.code, d.pathogen, d.description, d.id, dt.description FROM disease d
                                       INNER JOIN disease_type dt ON d.id = dt.id`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var d models.Disease
		err := rows.Scan(&d.Code, &d.Pathogen, &d.Description, &d.DiseaseTypeID, &d.DiseaseType)
		if err != nil {
			return nil, err
		}

		res = append(res, d)
	}

	return res, nil
}

func (r *diseaseRepository) CreateDiseaseType(ctx context.Context, dt *models.DiseaseType) (id int, err error) {
	q := `INSERT INTO disease_type (description) VALUES ($1) RETURNING id`

	err = r.db.QueryRow(ctx, q, dt.Description).Scan(&id)
	if err != nil {
		return
	}

	return
}

func (r *diseaseRepository) UpdateDiseaseType(ctx context.Context, dt *models.DiseaseType) (err error) {
	q := `UPDATE disease_type SET description = $1 WHERE id = $2`

	res, err := r.db.Exec(ctx, q, dt.Description, dt.ID)
	if err != nil {
		return
	}

	if count := res.RowsAffected(); count == 0 {
		return fmt.Errorf("no disease type with id %v", dt.ID)
	}

	return
}

func (r *diseaseRepository) DeleteDiseaseType(ctx context.Context, id int) (err error) {
	q := `DELETE FROM disease_type WHERE id = $1`

	res, err := r.db.Exec(ctx, q, id)
	if err != nil {
		return
	}

	if count := res.RowsAffected(); count == 0 {
		return fmt.Errorf("no disease type with id %v", id)
	}

	return
}

func (r *diseaseRepository) GetDiseaseType(ctx context.Context, id int) (dt *models.DiseaseType, err error) {
	q := `SELECT * FROM disease_type WHERE id = $1`

	dt = new(models.DiseaseType)

	err = r.db.QueryRow(ctx, q, id).Scan(&dt.ID, &dt.Description)
	if err != nil {
		return
	}

	return
}

func (r *diseaseRepository) GetAllDiseaseTypes(ctx context.Context) ([]models.DiseaseType, error) {
	var res []models.DiseaseType

	q := `SELECT * FROM disease_type`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var dt models.DiseaseType
		err := rows.Scan(&dt.ID, &dt.Description)
		if err != nil {
			return nil, err
		}

		res = append(res, dt)
	}

	return res, nil
}

func NewDiseaseRepository(db *pgxpool.Pool, cfg config.Postgres) DiseaseRepository {
	return &diseaseRepository{
		db:  db,
		cfg: cfg,
	}
}
