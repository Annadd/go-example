package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func ListOperation(cli *redis.Client) {
	var key = "fruit" //list key

	cli.RPush(key, "apple")      //list tail add element
	cli.LPush(key, "banana")     //list head add element
	cli.LPush(key, "watermelon") //list head add element

	len, err := cli.LLen(key).Result() //get list length
	if err != nil {
		panic(err)
	}
	fmt.Println(key+" len: ", len)

	ele, err := cli.RPop(key).Result() //remove tail element and return tail value
	if err != nil {
		panic(err)
	}
	fmt.Println(key, ele)
}
