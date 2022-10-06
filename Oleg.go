package lesson3

import (
	"fmt"
)

func fib(n int, carta map[int]int) int {

	if val, ok := carta[n]; ok {
		return val
	}
	carta[n] = fib(n-1, carta) + fib(n-2, carta)

	return carta[n]

}

func main() {
	var n int
	carta := map[int]int{

		0: 0,
		1: 1,
	}
	fmt.Println("Ведите значение: ")
	fmt.Scanln(&n)
	fmt.Println("Результат: ", fib(n, carta))
	fmt.Println(carta)
}
