package service

import (
	"encoding/json"
	"hte-device-update-dispatcher/internal/domain"
	"hte-device-update-dispatcher/internal/repository"
)

type DeviceService interface {
	Update(payload domain.Payload) error
}

type deviceService struct {
	repo repository.DeviceRepository
}

func NewDeviceService(repo repository.DeviceRepository) DeviceService {
	return &deviceService{repo: repo}
}

func (s *deviceService) Update(payload domain.Payload) error {
	payloadbytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return s.repo.Update(string(payloadbytes))
}
