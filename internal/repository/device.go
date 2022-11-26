package repository

import "log"

type DeviceRepository interface {
	Update(payload string) error
}

type deviceRepository struct {
}

func NewDeviceRepository() DeviceRepository {
	return &deviceRepository{}
}

func (r *deviceRepository) Update(payload string) error {
	log.Println(payload)
	return nil
}
