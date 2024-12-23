package redistest

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/redis"
)

func StartContainer(ctx context.Context) (testcontainers.Container, string, error) {
	redisContainer, err := redis.Run(ctx, "redis:7.4-alpine")
	if err != nil {
		return nil, "", err
	}

	connStr, err := redisContainer.ConnectionString(ctx)
	// remove redis:// from connstr
	connStr = connStr[8:]
	if err != nil {
		return nil, "", err
	}
	return redisContainer, connStr, nil
}
