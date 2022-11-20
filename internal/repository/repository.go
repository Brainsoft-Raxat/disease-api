package repository

import (
	"context"
	"disease-api/internal/app/config"
	"disease-api/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	DiseaseRepository
	DiscoverRepository
	DoctorRepository
	CountryRepository
	PublicServantRepository
	RecordRepository
}

type DiseaseRepository interface {
	// DiseaseType
	CreateDiseaseType(ctx context.Context, dt *models.DiseaseType) (int, error)
	UpdateDiseaseType(ctx context.Context, dt *models.DiseaseType) error
	DeleteDiseaseType(ctx context.Context, id int) error
	GetDiseaseType(ctx context.Context, id int) (*models.DiseaseType, error)
	GetAllDiseaseTypes(ctx context.Context) ([]models.DiseaseType, error)

	// Disease
	CreateDisease(ctx context.Context, d *models.Disease) (string, error)
	UpdateDisease(ctx context.Context, d *models.Disease) error
	DeleteDisease(ctx context.Context, code string) error
	GetDisease(ctx context.Context, code string) (*models.Disease, error)
	GetAllDiseases(ctx context.Context) ([]models.Disease, error)
}

type DiscoverRepository interface {
	CreateDiscover(ctx context.Context, d *models.Discover) (string, error)
	UpdateDiscover(ctx context.Context, d *models.Discover) error
	DeleteDiscover(ctx context.Context, country, diseaseCode string) error
	GetDiscover(ctx context.Context, country, diseaseCode string) (*models.Discover, error)
	GetAllDiscovers(ctx context.Context) ([]models.Discover, error)
}

type DoctorRepository interface {
	CreateDoctor(ctx context.Context, d *models.Doctor) (string, error)
	UpdateDoctor(ctx context.Context, d *models.Doctor) error
	DeleteDoctor(ctx context.Context, email string) error
	GetDoctor(ctx context.Context, email string) (*models.Doctor, error)
	GetAllDoctors(ctx context.Context) ([]models.Doctor, error)
}

type CountryRepository interface {
	CreateCountry(ctx context.Context, c *models.Country) (string, error)
	UpdateCountry(ctx context.Context, c *models.Country) error
	DeleteCountry(ctx context.Context, cname string) error
	GetCountry(ctx context.Context, cname string) (*models.Country, error)
	GetAllCountries(ctx context.Context) ([]models.Country, error)
}

type PublicServantRepository interface {
	CreatePublicServant(ctx context.Context, ps *models.PublicServant) (string, error)
	UpdatePublicServant(ctx context.Context, ps *models.PublicServant) error
	DeletePublicServant(ctx context.Context, email string) error
	GetPublicServant(ctx context.Context, email string) (*models.PublicServant, error)
	GetAllPublicServants(ctx context.Context) ([]models.PublicServant, error)
}

type RecordRepository interface {
	CreateRecord(ctx context.Context, r *models.Record) (string, error)
	UpdateRecord(ctx context.Context, r *models.Record) error
	DeleteRecord(ctx context.Context, email, cname, diseaseCode string) error
	GetRecordsFilter(ctx context.Context, email, cname, diseaseCode string) ([]models.Record, error)
}

func New(db *pgxpool.Pool, cfg *config.Config) *Repository {
	return &Repository{
		NewDiseaseRepository(db, cfg.Postgres),
		NewDiscoverRepository(db, cfg.Postgres),
		NewDoctorRepository(db, cfg.Postgres),
		NewCountryRepository(db, cfg.Postgres),
		NewPublicServantRepository(db, cfg.Postgres),
		NewRecordRepository(db, cfg.Postgres),
	}
}
