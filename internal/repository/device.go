package repository

import (
	"context"
	"hte-device-update-dispatcher/internal/defines"

	"github.com/go-redis/redis/v9"
)

type DeviceRepository interface {
	Update(payload string) error
}

type deviceRepository struct {
	redisClient *redis.Client
}

func NewDeviceRepository(redisClient *redis.Client) DeviceRepository {
	return &deviceRepository{redisClient: redisClient}
}

func (r *deviceRepository) Update(payload string) error {
	ctx := context.Background()
	return r.redisClient.RPush(ctx, defines.QueueDeviceUpdate, payload).Err()
}
