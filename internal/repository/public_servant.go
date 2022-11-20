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

type publicServantRepository struct {
	db  *pgxpool.Pool
	cfg config.Postgres
}

func (r *publicServantRepository) CreatePublicServant(ctx context.Context, ps *models.PublicServant) (string, error) {
	var email string

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return "", err
	}

	q := `INSERT INTO users (email, name, surname, salary, phone, cname) VALUES ($1, $2, $3, $4, $5, $6) RETURNING email`
	dRow, err := tx.Query(ctx, q, ps.Email, ps.Name, ps.Surname, ps.Salary, ps.Phone, ps.CName)
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
			return "", fmt.Errorf("error occurred while scanning public_servant in users: %v", err)
		}
	}
	dRow.Close()

	q = `INSERT INTO public_servant (email, department) VALUES ($1, $2) RETURNING email`
	dRow, err = tx.Query(ctx, q, ps.Email, ps.Department)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return "", fmt.Errorf("error occurred while creaing public_servant in public_servant: %v", err)
	}
	if dRow.Next() {
		err = dRow.Scan(&email)
		if err != nil {
			errTX := tx.Rollback(ctx)
			if errTX != nil {
				log.Printf("ERROR: transaction: %s", errTX)
			}
			return "", fmt.Errorf("error occurred while creating public_servant in public_servant: %v", err)
		}
	}
	dRow.Close()

	err = tx.Commit(ctx)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction error: %s", errTX)
		}
		return "", fmt.Errorf("error occurred while creating public_servant: %v", err)
	}

	return email, nil
}

func (r *publicServantRepository) UpdatePublicServant(ctx context.Context, ps *models.PublicServant) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	q := `UPDATE users SET email = $1, cname = $2, name = $3, surname = $4, phone = $5, salary = $6 WHERE email = $7`
	_, err = tx.Exec(ctx, q, ps.Email, ps.CName, ps.Name, ps.Surname, ps.Phone, ps.Salary, ps.Email)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return fmt.Errorf("error occurred while updating public_servant info in users table: %v", err)
	}

	q = `UPDATE public_servant SET email = $1, department = $2 WHERE email = $3`
	_, err = tx.Exec(ctx, q, ps.Email, ps.Department, ps.Email)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return fmt.Errorf("error occurred while updating public_servant INFO in public_servant: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction error: %s", errTX)
		}
		return fmt.Errorf("error occurred while deleting public_servant from users: %v", err)
	}
	return nil
}

func (r *publicServantRepository) DeletePublicServant(ctx context.Context, email string) error {
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
		return fmt.Errorf("error occurred while deleting public_servant from doctors: %v", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction error: %s", errTX)
		}
		return fmt.Errorf("error occurred while deleting public_servant from users: %v", err)
	}

	return nil
}

func (r *publicServantRepository) GetPublicServant(ctx context.Context, email string) (ps *models.PublicServant, err error) {
	q := `SELECT u.email, u.name, u.surname, u.phone, u.salary, u.cname, ps.department  FROM users u 
    INNER JOIN public_servant ps on u.email = ps.email WHERE u.email = $1`

	ps = new(models.PublicServant)

	err = r.db.QueryRow(ctx, q, email).Scan(&ps.Email, &ps.Name, &ps.Surname, &ps.Phone, &ps.Salary, &ps.CName, &ps.Department)
	if err != nil {
		return
	}

	return
}

func (r *publicServantRepository) GetAllPublicServants(ctx context.Context) ([]models.PublicServant, error) {
	var res []models.PublicServant

	q := `SELECT u.email, u.name, u.surname, u.phone, u.salary, u.cname, ps.department  FROM users u 
    INNER JOIN public_servant ps on u.email = ps.email`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ps models.PublicServant
		err := rows.Scan(&ps.Email, &ps.Name, &ps.Surname, &ps.Phone, &ps.Salary, &ps.CName, &ps.Department)
		if err != nil {
			return nil, err
		}

		res = append(res, ps)
	}

	return res, nil
}

func NewPublicServantRepository(db *pgxpool.Pool, cfg config.Postgres) PublicServantRepository {
	return &publicServantRepository{
		db:  db,
		cfg: cfg,
	}
}
