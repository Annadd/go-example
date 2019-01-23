package main

import (
	"fmt"
	"time"
	"github.com/tarm/goserial"
)

func main() {
	config := &serial.Config{
		Name:        "COM1",
		Baud:        115200,
		ReadTimeout: time.Second * 2,
	}

	io, err := serial.OpenPort(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := io.Write([]byte("cc"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}