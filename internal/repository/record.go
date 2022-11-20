package repository

import (
	"context"
	"disease-api/internal/app/config"
	"disease-api/internal/models"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type recordRepository struct {
	db  *pgxpool.Pool
	cfg config.Postgres
}

func (r *recordRepository) CreateRecord(ctx context.Context, rec *models.Record) (email string, err error) {
	q := `INSERT INTO record (email, cname, disease_code, total_deaths, total_patients) VALUES ($1, $2, $3, $4, $5) RETURNING email`

	err = r.db.QueryRow(ctx, q, rec.Email, rec.CName, rec.DiseaseCode, rec.TotalDeaths, rec.TotalPatients).Scan(&email)
	if err != nil {
		return
	}

	return
}

func (r *recordRepository) UpdateRecord(ctx context.Context, rec *models.Record) error {
	q := `UPDATE record SET email = $1, cname = $2, disease_code = $3, total_deaths = $4::int, total_patients = $5::int WHERE email = $6 AND cname = $7 AND disease_code = $8`

	res, err := r.db.Exec(ctx, q, rec.Email, rec.CName, rec.DiseaseCode, rec.TotalDeaths, rec.TotalPatients, rec.Email, rec.TotalDeaths, rec.DiseaseCode)
	if err != nil {
		return err
	}

	if count := res.RowsAffected(); count == 0 {
		return fmt.Errorf("no record with cname %v", rec.Email)
	}

	return nil

}

func (r *recordRepository) DeleteRecord(ctx context.Context, email, cname, diseaseCode string) error {
	q := `DELETE FROM record WHERE email = $1 AND cname = $2 AND disease_code = $3`

	res, err := r.db.Exec(ctx, q, email, cname, diseaseCode)
	if err != nil {
		return err
	}

	if count := res.RowsAffected(); count == 0 {
		return fmt.Errorf("no country with cname %v", cname)
	}

	return nil
}

func (r *recordRepository) GetRecordsFilter(ctx context.Context, email, cname, diseaseCode string) ([]models.Record, error) {
	var res []models.Record

	q := `SELECT email, cname, disease_code, total_deaths, total_patients FROM record `

	var args []interface{}

	if !(email == "" && cname == "") || diseaseCode != "" {
		q += `WHERE `

		flag := false
		i := 1

		if email != "" {
			flag = true
			q += fmt.Sprintf("email = $%v ", i)
			i++
			args = append(args, email)
		}

		if cname != "" {
			if flag {
				q += fmt.Sprintf("AND cname = $%v ", i)
			} else {
				q += fmt.Sprintf("cname = $%v ", i)
			}
			i++
			args = append(args, cname)
		}

		if diseaseCode != "" {
			if flag {
				q += fmt.Sprintf("AND disease_code = $%v ", i)
			} else {
				q += fmt.Sprintf("disease_code = $%v ", i)
			}
			i++
			args = append(args, diseaseCode)
		}
	}

	rows, err := r.db.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var rec models.Record
		err := rows.Scan(&rec.Email, &rec.CName, &rec.DiseaseCode, &rec.TotalDeaths, &rec.TotalPatients)
		if err != nil {
			return nil, err
		}

		res = append(res, rec)
	}

	return res, nil
}

func NewRecordRepository(db *pgxpool.Pool, cfg config.Postgres) RecordRepository {
	return &recordRepository{
		db:  db,
		cfg: cfg,
	}
}
