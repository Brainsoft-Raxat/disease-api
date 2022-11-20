package service

import (
	"context"
	"disease-api/internal/models"
	"disease-api/internal/repository"
	"time"
)

type discoverService struct {
	discoverRepo repository.DiscoverRepository
}

func (s *discoverService) CreateDiscover(ctx context.Context, d *models.Discover) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	code, err := s.discoverRepo.CreateDiscover(ctx, d)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (s *discoverService) UpdateDiscover(ctx context.Context, d *models.Discover) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	err := s.discoverRepo.UpdateDiscover(ctx, d)
	if err != nil {
		return err
	}

	return nil
}

func (s *discoverService) DeleteDiscover(ctx context.Context, country, diseaseCode string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	err := s.discoverRepo.DeleteDiscover(ctx, country, diseaseCode)
	if err != nil {
		return err
	}

	return nil
}

func (s *discoverService) GetDiscover(ctx context.Context, country, diseaseCode string) (*models.Discover, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	d, err := s.discoverRepo.GetDiscover(ctx, country, diseaseCode)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (s *discoverService) GetAllDiscovers(ctx context.Context) ([]models.Discover, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	res, err := s.discoverRepo.GetAllDiscovers(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewDiscoverService(repo *repository.Repository) DiscoverService {
	return &discoverService{discoverRepo: repo.DiscoverRepository}
}
