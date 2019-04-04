package main

import(
	"fmt"
	"os"
)

func main(){
	process_name := os.Args[0]
	fmt.Println(process_name)
}
