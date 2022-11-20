package service

import (
	"context"
	"disease-api/internal/models"
	"disease-api/internal/repository"
	"time"
)

type doctorService struct {
	doctorRepo repository.DoctorRepository
}

func (s *doctorService) CreateDoctor(ctx context.Context, d *models.Doctor) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	email, err := s.doctorRepo.CreateDoctor(ctx, d)
	if err != nil {
		return "", err
	}

	return email, nil
}

func (s *doctorService) UpdateDoctor(ctx context.Context, d *models.Doctor) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	err := s.doctorRepo.UpdateDoctor(ctx, d)
	if err != nil {
		return err
	}

	return nil
}

func (s *doctorService) DeleteDoctor(ctx context.Context, email string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	err := s.doctorRepo.DeleteDoctor(ctx, email)
	if err != nil {
		return err
	}

	return nil
}

func (s *doctorService) GetDoctor(ctx context.Context, email string) (*models.Doctor, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	d, err := s.doctorRepo.GetDoctor(ctx, email)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (s *doctorService) GetAllDoctors(ctx context.Context) ([]models.Doctor, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	resp, err := s.doctorRepo.GetAllDoctors(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewDoctorService(repo *repository.Repository) DoctorService {
	return &doctorService{doctorRepo: repo.DoctorRepository}
}
