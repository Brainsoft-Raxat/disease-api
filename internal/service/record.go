package service

import (
	"context"
	"disease-api/internal/models"
	"disease-api/internal/repository"
	"time"
)

type recordService struct {
	recordRepo repository.RecordRepository
}

func (s *recordService) CreateRecord(ctx context.Context, r *models.Record) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	email, err := s.recordRepo.CreateRecord(ctx, r)
	if err != nil {
		return "", err
	}

	return email, nil
}

func (s *recordService) UpdateRecord(ctx context.Context, r *models.Record) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	err := s.recordRepo.UpdateRecord(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *recordService) DeleteRecord(ctx context.Context, email, cname, diseaseCode string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	err := s.recordRepo.DeleteRecord(ctx, email, cname, diseaseCode)
	if err != nil {
		return err
	}

	return nil
}

func (s *recordService) GetRecordsFilter(ctx context.Context, email, cname, diseaseCode string) ([]models.Record, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	resp, err := s.recordRepo.GetRecordsFilter(ctx, email, cname, diseaseCode)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewRecordService(repo *repository.Repository) RecordService {
	return &recordService{recordRepo: repo}
}
