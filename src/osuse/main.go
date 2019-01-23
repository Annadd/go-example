package main 

import(
	"fmt"
	"os"
)


func main(){
	for idx, args := range os.Args{
		fmt.Println(idx,args)
	} 
}