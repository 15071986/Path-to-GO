package main

import (
	"fmt"
	"sync"
)

var v int64
var rw sync.RWMutex

func main() {

	for i := 0; i < 90; i++ {
		go func() {
			rw.RLock()
			defer rw.RUnlock()
			if v%10 == 0 {
				fmt.Println("ok")
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				rw.Lock()
				v++
				rw.Unlock()
			}
		}()
	}
}
