package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func SetOperation(cli *redis.Client) {
	var key = "setcollection"

	//add element
	cli.SAdd(key, "obama")
	cli.SAdd(key, "hillary")
	cli.SAdd(key, "the elder")
	cli.SAdd(key, "pass")
	cli.SAdd(key, "pass")

	isMember, err := cli.SIsMember(key, "bush").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("is push in "+key+":", isMember)

	all, err := cli.SMembers(key).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("all member: ", all)
}
