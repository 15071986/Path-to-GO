package fib

import (
	"fmt"
)

func Fib(n int, carta map[int]int) int {

	if val, ok := carta[n]; ok {
		return val
	}
	carta[n] = Fib(n-1, carta) + Fib(n-2, carta)

	return carta[n]

}

func Main() {
	var n int
	carta := map[int]int{

		0: 0,
		1: 1,
	}
	fmt.Println("enter a number to calculate: ")
	fmt.Scanln(&n)
	fmt.Println("Fibonachi number is: ", Fib(n, carta))
	fmt.Println(carta)
}
