package main

/*
#include "util.c"
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.sum(10, 20))
}
