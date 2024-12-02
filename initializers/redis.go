package initializers

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func RedisInitializer() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // Use default DB
		Protocol: 2, // Connection protocol
	})
}
