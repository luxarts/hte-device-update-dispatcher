package domain

import "hte-device-update-dispatcher/internal/domain/pb"

type Coordinates struct {
	Latitude  *float64 `json:"lat"`
	Longitude *float64 `json:"lon"`
}

type Payload struct {
	DeviceID    string      `json:"device_id"`
	Timestamp   int64       `json:"ts"`
	Battery     int64       `json:"bat"`
	Coordinates Coordinates `json:"coords"`
}

func (p *Payload) IsValid() bool {
	return p.DeviceID != "" && p.Timestamp > 0 && p.Battery >= 0 && p.Coordinates.Latitude != nil && p.Coordinates.Longitude != nil
}

func (p *Payload) ToProto() *pb.Payload {
	return &pb.Payload{
		DeviceID:  p.DeviceID,
		Timestamp: p.Timestamp,
		Battery:   p.Battery,
		Coordinates: &pb.Payload_Coordinates{
			Latitude:  float32(*p.Coordinates.Latitude),
			Longitude: float32(*p.Coordinates.Longitude),
		},
	}
}
