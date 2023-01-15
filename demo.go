package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int, 10)
	go func() {
		ch <- 1
		time.Sleep(3 * time.Second)
		close(ch)
	}()
	go func() {
		ch <- 2
		time.Sleep(2 * time.Second)
		close(ch)
	}()
	go func() {
		ch <- 3
		time.Sleep(1 * time.Second)
		close(ch)
	}()

	val := <-ch
	fmt.Println("Канал №1 закрыт!", val)
	val1 := <-ch
	fmt.Println("Канал №2 закрыт!", val1)
	val2 := <-ch
	fmt.Println("Канал №3 закрыт!", val2)

}
