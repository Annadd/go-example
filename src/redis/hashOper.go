package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func HashOperation(cli *redis.Client) {
	var key = "user"
	cli.HSet(key, "name", "annadd")
	cli.HSet(key, "age", "18")
	cli.HSet(key, "sex", "man")

	len, err := cli.HLen(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("hash len", len)

	val, err := cli.HGetAll(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
