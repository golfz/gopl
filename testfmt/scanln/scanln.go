package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i:= 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("count:", i)
		}
	}()

	fmt.Scanln()
}
