package service

import (
	"encoding/json"
	"hte-device-update-dispatcher/internal/domain"
	"hte-device-update-dispatcher/internal/repository"
	"log"

	"google.golang.org/protobuf/proto"
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
	pbp := payload.ToProto()
	payloadbytes, err := proto.Marshal(pbp)
	if err != nil {
		return err
	}

	pbJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	log.Println(pbJson)
	return s.repo.Update(string(payloadbytes))
}
