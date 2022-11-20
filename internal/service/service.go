package service

import (
	"context"
	"disease-api/internal/models"
	"disease-api/internal/repository"
)

type DiseaseService interface {
	CreateDisease(ctx context.Context, d *models.Disease) (string, error)
	UpdateDisease(ctx context.Context, d *models.Disease) error
	DeleteDisease(ctx context.Context, code string) error
	GetDisease(ctx context.Context, code string) (*models.Disease, error)
	GetAllDiseases(ctx context.Context) ([]models.Disease, error)

	CreateDiseaseType(ctx context.Context, dt *models.DiseaseType) (int, error)
	UpdateDiseaseType(ctx context.Context, dt *models.DiseaseType) error
	DeleteDiseaseType(ctx context.Context, id int) error
	GetDiseaseType(ctx context.Context, id int) (*models.DiseaseType, error)
	GetAllDiseaseTypes(ctx context.Context) ([]models.DiseaseType, error)
}

type DiscoverService interface {
	CreateDiscover(ctx context.Context, d *models.Discover) (string, error)
	UpdateDiscover(ctx context.Context, d *models.Discover) error
	DeleteDiscover(ctx context.Context, country, diseaseCode string) error
	GetDiscover(ctx context.Context, country, diseaseCode string) (*models.Discover, error)
	GetAllDiscovers(ctx context.Context) ([]models.Discover, error)
}

type DoctorService interface {
	CreateDoctor(ctx context.Context, d *models.Doctor) (string, error)
	UpdateDoctor(ctx context.Context, d *models.Doctor) error
	DeleteDoctor(ctx context.Context, email string) error
	GetDoctor(ctx context.Context, email string) (*models.Doctor, error)
	GetAllDoctors(ctx context.Context) ([]models.Doctor, error)
}

type CountryService interface {
	CreateCountry(ctx context.Context, c *models.Country) (string, error)
	UpdateCountry(ctx context.Context, c *models.Country) error
	DeleteCountry(ctx context.Context, cname string) error
	GetCountry(ctx context.Context, cname string) (*models.Country, error)
	GetAllCountries(ctx context.Context) ([]models.Country, error)
}

type PublicServantService interface {
	CreatePublicServant(ctx context.Context, ps *models.PublicServant) (string, error)
	UpdatePublicServant(ctx context.Context, ps *models.PublicServant) error
	DeletePublicServant(ctx context.Context, email string) error
	GetPublicServant(ctx context.Context, email string) (*models.PublicServant, error)
	GetAllPublicServants(ctx context.Context) ([]models.PublicServant, error)
}

type RecordService interface {
	CreateRecord(ctx context.Context, r *models.Record) (string, error)
	UpdateRecord(ctx context.Context, r *models.Record) error
	DeleteRecord(ctx context.Context, email, cname, diseaseCode string) error
	GetRecordsFilter(ctx context.Context, email, cname, diseaseCode string) ([]models.Record, error)
}

type Service struct {
	DiseaseService
	DiscoverService
	DoctorService
	CountryService
	PublicServantService
	RecordService
}

func New(repos *repository.Repository) *Service {
	return &Service{
		DiseaseService:       NewDiseaseService(repos),
		DiscoverService:      NewDiscoverService(repos),
		DoctorService:        NewDoctorService(repos),
		CountryService:       NewCountryService(repos),
		PublicServantService: NewPublicServantService(repos),
		RecordService:        NewRecordService(repos),
	}
}
