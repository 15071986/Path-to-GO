package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 4; i++ {
		go threeflow()
	}
	time.Sleep(2 * time.Second)
}

func threeflow() {
	time.Sleep(1 * time.Second)
	fmt.Println("Good")
}
