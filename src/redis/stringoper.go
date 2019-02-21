package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func StringOperationTest(cli *redis.Client) {
	err := cli.Set("name", "Annadd", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := cli.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)

	err = cli.Set("age", "20", 1*time.Second).Err()
	if err != nil {
		panic(err)
	}

	cli.Incr("age")
	cli.Incr("age")
	cli.Decr("age")

	val, err = cli.Get("age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("age", val)

	time.Sleep(1 * time.Second)
	val, err = cli.Get("age").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("age", val)
}
