package service

import (
	"context"
	"disease-api/internal/models"
	"disease-api/internal/repository"
	"time"
)

type countryService struct {
	countryRepo repository.CountryRepository
}

func (s *countryService) CreateCountry(ctx context.Context, c *models.Country) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	cname, err := s.countryRepo.CreateCountry(ctx, c)
	if err != nil {
		return "", err
	}

	return cname, nil
}

func (s *countryService) UpdateCountry(ctx context.Context, c *models.Country) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	err := s.countryRepo.UpdateCountry(ctx, c)
	if err != nil {
		return err
	}

	return nil
}

func (s *countryService) DeleteCountry(ctx context.Context, cname string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	err := s.countryRepo.DeleteCountry(ctx, cname)
	if err != nil {
		return err
	}

	return nil
}

func (s *countryService) GetCountry(ctx context.Context, cname string) (*models.Country, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	c, err := s.countryRepo.GetCountry(ctx, cname)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *countryService) GetAllCountries(ctx context.Context) ([]models.Country, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	resp, err := s.countryRepo.GetAllCountries(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewCountryService(repo *repository.Repository) CountryService {
	return &countryService{countryRepo: repo.CountryRepository}
}
