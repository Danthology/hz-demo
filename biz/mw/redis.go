package mw

import "github.com/redis/go-redis/v9"

var RDS *redis.Client

func Init() {
	RDS = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
