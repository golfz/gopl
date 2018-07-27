package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for index, val := range os.Args[1:] {
		fmt.Println("(" + strconv.Itoa(index) + ")--> " + val)
	}
}
