package services

import (
	"go-xray-config/internal/repository"

	"github.com/google/uuid"
)

type XrayService interface {
	Add() (string, error)
	Delete(id uuid.UUID) error
}

type xrayService struct {
	xrayRepository repository.XrayRepository
	host           string
	publicKey      string
	serverName     string
}

func NewXrayService(xrayRepository repository.XrayRepository) XrayService {
	return &xrayService{
		xrayRepository: xrayRepository,
	}
}

func (s *xrayService) Add() (string, error) {
	result, err := s.xrayRepository.Add()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *xrayService) Delete(id uuid.UUID) error {
	return s.xrayRepository.Delete(id)
}
