package main

import (
	"fmt"
	"os"
)

/*
This echo program prints the index and value of each of its arguments.
*/
func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
