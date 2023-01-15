package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		fmt.Println("worker", id, "finished job", j)
		results <- j + 1
	}
}

func main() {
	jobs := make(chan int, 1000)
	results := make(chan int, 1000)

	for w := 1000; w <= 1000; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 1000; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 1000; a++ {
		<-results
	}
}
