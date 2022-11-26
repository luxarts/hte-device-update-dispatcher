package controller

import (
	"encoding/json"
	"hte-device-update-dispatcher/internal/domain"
	"hte-device-update-dispatcher/internal/service"
	"log"
)

type DispatcherController interface {
	Handle(msg []byte)
}

type dispatcherController struct {
	svc service.DeviceService
}

func NewDispatcherController(svc service.DeviceService) DispatcherController {
	return &dispatcherController{svc: svc}
}

func (c *dispatcherController) Handle(msg []byte) {
	var payload domain.Payload
	err := json.Unmarshal(msg, &payload)
	if err != nil {
		log.Println(err)
		return
	}
	if !payload.IsValid() {
		log.Println("Invalid Payload")
		return
	}
	err = c.svc.Update(payload)
	if err != nil {
		log.Println(err)
		return
	}
}
