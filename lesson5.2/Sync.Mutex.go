package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func main() {

	for i := 0; i < 4; i++ {
		go threeflow()
	}
	time.Sleep(2 * time.Second)
}

func threeflow() {
	mutex.Lock()
	defer mutex.Unlock()
	time.Sleep(1 * time.Second)
	fmt.Println("Good")
}
