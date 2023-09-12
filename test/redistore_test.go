package test

import (
	"fmt"
	"testing"

	"gopkg.in/boj/redistore.v1"
)

func TestRedistore(t *testing.T) {
	store, err := redistore.NewRediStore(10, "tcp", "127.0.0.1:6379", "123456", []byte("swag-redis"))
	// redis.NewStore()
	if err != nil {
		fmt.Println("error:", err)
	}
	defer store.Close()
}
