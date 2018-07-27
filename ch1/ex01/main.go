package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("comamnd : " + os.Args[0])

	fmt.Println("os.Args :")

	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

}
