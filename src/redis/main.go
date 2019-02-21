package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func gRedisNew(host string, pwd string) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pwd,
		DB:       0,
	})

	pong, err := cli.Ping().Result()
	if err == nil {
		return cli
	}

	fmt.Println(pong, err)
	return nil
}

func main() {
	var host = "192.168.1.241:6379"
	var pwd = "cdsw"

	cli := gRedisNew(host, pwd)

	StringOperationTest(cli)
	ListOperation(cli)
	SetOperation(cli)
	HashOperation(cli)
}
