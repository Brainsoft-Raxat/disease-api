package service

import (
	"context"
	"disease-api/internal/models"
	"disease-api/internal/repository"
	"time"
)

type diseaseService struct {
	diseaseRepo repository.DiseaseRepository
}

func (s *diseaseService) CreateDisease(ctx context.Context, d *models.Disease) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	code, err := s.diseaseRepo.CreateDisease(ctx, d)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (s *diseaseService) UpdateDisease(ctx context.Context, d *models.Disease) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	err := s.diseaseRepo.UpdateDisease(ctx, d)
	if err != nil {
		return err
	}

	return nil
}

func (s *diseaseService) DeleteDisease(ctx context.Context, code string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	err := s.diseaseRepo.DeleteDisease(ctx, code)
	if err != nil {
		return err
	}

	return nil
}

func (s *diseaseService) GetDisease(ctx context.Context, code string) (*models.Disease, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	d, err := s.diseaseRepo.GetDisease(ctx, code)
	if err != nil {
		return nil, err
	}

	return d, err
}

func (s *diseaseService) GetAllDiseases(ctx context.Context) ([]models.Disease, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	resp, err := s.diseaseRepo.GetAllDiseases(ctx)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (s *diseaseService) CreateDiseaseType(ctx context.Context, dt *models.DiseaseType) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	id, err := s.diseaseRepo.CreateDiseaseType(ctx, dt)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *diseaseService) UpdateDiseaseType(ctx context.Context, dt *models.DiseaseType) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	err := s.diseaseRepo.UpdateDiseaseType(ctx, dt)
	if err != nil {
		return err
	}

	return nil
}

func (s *diseaseService) DeleteDiseaseType(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	err := s.diseaseRepo.DeleteDiseaseType(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *diseaseService) GetDiseaseType(ctx context.Context, id int) (*models.DiseaseType, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	dt, err := s.diseaseRepo.GetDiseaseType(ctx, id)
	if err != nil {
		return nil, err
	}

	return dt, nil
}

func (s *diseaseService) GetAllDiseaseTypes(ctx context.Context) ([]models.DiseaseType, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	resp, err := s.diseaseRepo.GetAllDiseaseTypes(ctx)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func NewDiseaseService(repo *repository.Repository) DiseaseService {
	return &diseaseService{diseaseRepo: repo.DiseaseRepository}
}
