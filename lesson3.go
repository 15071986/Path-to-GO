package main

import (
	"fmt"
	"math"
	"os"
)

var n int

func fact(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	var res float64

	for {
		var a, b, err float64
		var op string
		if err != nil {
			fmt.Println("ошибка: %s\n", err.Error())

			continue

			switch op {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				if b != 0 {
					fmt.Println(a / b)
				} else {
					fmt.Println("Деление на ноль, ошибка!")
				}
				return
			case "^":
				res = math.Pow(a, b)
			case "f":
				res = float64(fact(int(a)))

			default:
				fmt.Println("Операция выбрана неверно")
				os.Exit(1)
			}

			fmt.Printf("Результат выполнения операции: %.2f\n", res)
		}

	}
}
