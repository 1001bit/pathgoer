package refresh

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const expiration = 30 * 24 * time.Hour

type Storage struct {
	redisClient *redis.Client
}

func NewStorage(host, port string) *Storage {
	connStr := fmt.Sprintf("%s:%s", host, port)
	log.Println("connecting to redis on", connStr)

	return &Storage{
		redisClient: redis.NewClient(&redis.Options{
			Addr: connStr,
		}),
	}
}

func (r *Storage) GetUserIDAndRefresh(ctx context.Context, uuid string) (string, string, error) {
	userID, err := r.redisClient.Get(ctx, "uuid:"+uuid).Result()
	if err != nil {
		return "", "", err
	}
	// rotate token
	err = r.redisClient.Del(ctx, "uuid:"+uuid).Err()
	if err != nil {
		return "", "", err
	}

	newUUID, err := r.GenerateUUID(ctx, userID)
	return userID, newUUID, err
}

func (r *Storage) GenerateUUID(ctx context.Context, userID string) (string, error) {
	uuid := uuid.NewString()

	return uuid, r.redisClient.Set(ctx, "uuid:"+uuid, userID, expiration).Err()
}

func (r *Storage) DeleteUUID(ctx context.Context, uuid string) error {
	return r.redisClient.Del(ctx, "uuid:"+uuid).Err()
}