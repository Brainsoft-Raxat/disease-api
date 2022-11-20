package repository

import (
	"context"
	"disease-api/internal/app/config"
	"disease-api/internal/models"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type doctorRepository struct {
	db  *pgxpool.Pool
	cfg config.Postgres
}

func (r *doctorRepository) CreateDoctor(ctx context.Context, d *models.Doctor) (string, error) {
	var email string

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return "", err
	}

	q := `INSERT INTO users (email, name, surname, salary, phone, cname) VALUES ($1, $2, $3, $4, $5, $6) RETURNING email`
	dRow, err := tx.Query(ctx, q, d.Email, d.Name, d.Surname, d.Salary, d.Phone, d.CName)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			return "", errTX
		}

		return "", err
	}
	if dRow.Next() {
		err = dRow.Scan(&email)
		if err != nil {
			errTX := tx.Rollback(ctx)
			if errTX != nil {
				log.Printf("ERROR: transaction: %s", errTX)
			}
			return "", fmt.Errorf("error occurred while scanning doctor in users: %v", err)
		}
	}
	dRow.Close()

	q = `INSERT INTO doctor (email, degree) VALUES ($1, $2) RETURNING email`
	dRow, err = tx.Query(ctx, q, d.Email, d.Degree)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return "", fmt.Errorf("error occurred while creaing doctors in doctors: %v", err)
	}
	if dRow.Next() {
		err = dRow.Scan(&email)
		if err != nil {
			errTX := tx.Rollback(ctx)
			if errTX != nil {
				log.Printf("ERROR: transaction: %s", errTX)
			}
			return "", fmt.Errorf("error occurred while creating doctors in doctors: %v", err)
		}
	}
	dRow.Close()

	q = `INSERT INTO specialize (id, email) VALUES ($1, $2) RETURNING email`
	dRow, err = tx.Query(ctx, q, d.DiseaseTypeID, d.Email)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return "", fmt.Errorf("error occurred while creaing specialize in specialize: %v", err)
	}
	if dRow.Next() {
		err = dRow.Scan(&email)
		if err != nil {
			errTX := tx.Rollback(ctx)
			if errTX != nil {
				log.Printf("ERROR: transaction: %s", errTX)
			}
			return "", fmt.Errorf("error occurred while creating specialize in specialize: %v", err)
		}
	}
	dRow.Close()

	err = tx.Commit(ctx)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction error: %s", errTX)
		}
		return "", fmt.Errorf("error occurred while creating doctor: %v", err)
	}

	return email, nil
}

func (r *doctorRepository) UpdateDoctor(ctx context.Context, d *models.Doctor) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	q := `UPDATE users SET email = $1, cname = $2, name = $3, surname = $4, phone = $5, salary = $6 WHERE email = $7`
	_, err = tx.Exec(ctx, q, d.Email, d.CName, d.Name, d.Surname, d.Phone, d.Salary, d.Email)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return fmt.Errorf("error occurred while updating doctor info in users table: %v", err)
	}

	q = `UPDATE doctor SET email = $1, degree = $2 WHERE email = $3`
	_, err = tx.Exec(ctx, q, d.Email, d.Degree, d.Email)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return fmt.Errorf("error occurred while updating doctor INFO in doctors: %v", err)
	}

	q = `UPDATE specialize SET email = $1, id = $2 WHERE email = $3`
	_, err = tx.Exec(ctx, q, d.Email, d.DiseaseTypeID, d.Email)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return fmt.Errorf("error occurred while updating specialize INFO in specialize: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction error: %s", errTX)
		}
		return fmt.Errorf("error occurred while deleting doctor from users: %v", err)
	}
	return nil
}

func (r *doctorRepository) DeleteDoctor(ctx context.Context, email string) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	q := `DELETE FROM users WHERE email = $1`
	_, err = tx.Exec(ctx, q, email)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction error: %s", errTX)
		}
		return fmt.Errorf("error occurred while deleting doctor from doctors: %v", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction error: %s", errTX)
		}
		return fmt.Errorf("error occurred while deleting doctor from users: %v", err)
	}

	return nil
}

func (r *doctorRepository) GetDoctor(ctx context.Context, email string) (d *models.Doctor, err error) {
	q := `SELECT u.email, u.name, u.surname, u.phone, u.salary, u.cname, d.degree, dt.id, dt.description FROM users u 
    INNER JOIN doctor d on u.email = d.email
    INNER JOIN specialize s on d.email = s.email
    INNER JOIN disease_type dt on dt.id = s.id WHERE u.email = $1`

	d = new(models.Doctor)

	err = r.db.QueryRow(ctx, q, email).Scan(&d.Email, &d.Name, &d.Surname, &d.Phone, &d.Salary, &d.CName, &d.Degree, &d.DiseaseTypeID, &d.DiseaseType)
	if err != nil {
		return
	}

	return
}

func (r *doctorRepository) GetAllDoctors(ctx context.Context) ([]models.Doctor, error) {
	var res []models.Doctor

	q := `SELECT u.email, u.name, u.surname, u.phone, u.salary, u.cname, d.degree, dt.id, dt.description FROM users u 
    INNER JOIN doctor d on u.email = d.email
    INNER JOIN specialize s on d.email = s.email
    INNER JOIN disease_type dt on dt.id = s.id`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var d models.Doctor
		err := rows.Scan(&d.Email, &d.Name, &d.Surname, &d.Phone, &d.Salary, &d.CName, &d.Degree, &d.DiseaseTypeID, &d.DiseaseType)
		if err != nil {
			return nil, err
		}

		res = append(res, d)
	}

	return res, nil
}

func NewDoctorRepository(db *pgxpool.Pool, cfg config.Postgres) DoctorRepository {
	return &doctorRepository{
		db:  db,
		cfg: cfg,
	}
}
