package otp

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const expiration = 5 * time.Minute

type Storage struct {
	redisClient *redis.Client
}

func NewStorage(host, port string) *Storage {
	return &Storage{
		redisClient: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", host, port),
		}),
	}
}

func (r *Storage) VerifyOTP(ctx context.Context, email, otp string) bool {
	resOtp, err := r.redisClient.Get(ctx, "email:"+email).Result()
	if err != nil || resOtp != otp {
		return false
	}

	r.redisClient.Del(ctx, email).Err()
	return true
}

func (r *Storage) GenerateOTP(ctx context.Context, email string) (string, error) {
	otp, err := generateOTP()
	if err != nil {
		return "", err
	}

	return otp, r.redisClient.Set(ctx, "email:"+email, otp, expiration).Err()
}
