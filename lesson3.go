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
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число (для расчета факториала введите 0): ")
	fmt.Scanln(&b)

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, f): ")
	fmt.Scanln(&op)

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
