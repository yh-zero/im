package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()

func TestSet(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Set(ctx, "key2", "value222222", time.Second*5).Err()
	if err != nil {
		panic(err)
	}
}

func TestGet(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	val, err := rdb.Get(ctx, "key2 =========》").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1", val)
}
