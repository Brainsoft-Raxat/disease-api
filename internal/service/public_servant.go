package service

import (
	"context"
	"disease-api/internal/models"
	"disease-api/internal/repository"
	"time"
)

type publicServantService struct {
	publicServantRepo repository.PublicServantRepository
}

func (s *publicServantService) CreatePublicServant(ctx context.Context, ps *models.PublicServant) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	email, err := s.publicServantRepo.CreatePublicServant(ctx, ps)
	if err != nil {
		return "", err
	}

	return email, nil
}

func (s *publicServantService) UpdatePublicServant(ctx context.Context, ps *models.PublicServant) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	err := s.publicServantRepo.UpdatePublicServant(ctx, ps)
	if err != nil {
		return err
	}

	return nil
}

func (s *publicServantService) DeletePublicServant(ctx context.Context, email string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	err := s.publicServantRepo.DeletePublicServant(ctx, email)
	if err != nil {
		return err
	}

	return nil
}

func (s *publicServantService) GetPublicServant(ctx context.Context, email string) (*models.PublicServant, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	d, err := s.publicServantRepo.GetPublicServant(ctx, email)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (s *publicServantService) GetAllPublicServants(ctx context.Context) ([]models.PublicServant, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	resp, err := s.publicServantRepo.GetAllPublicServants(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewPublicServantService(repo *repository.Repository) PublicServantService {
	return &publicServantService{publicServantRepo: repo.PublicServantRepository}
}
