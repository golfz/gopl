package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println()

	start1 := time.Now()
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
	fmt.Printf("[-1-] for : %.3f microsecs elapsed\n", time.Since(start1).Seconds()*1000000)

	fmt.Println()

	start2 := time.Now()
	for index, val := range os.Args[1:] {
		fmt.Println("(" + strconv.Itoa(index) + ")--> " + val)
	}
	fmt.Printf("[-2-] for range : %.3f microsecs elapsed\n", time.Since(start2).Seconds()*1000000)

	fmt.Println()

	start3 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("[-3-] strings.Join : %.3f microsecs elapsed\n", time.Since(start3).Seconds()*1000000)

	fmt.Println()
}
