package main

import (
	"fmt"
	"math"
)

func main() {

	var S float64
	var L float64
	var D float64

	fmt.Println("Введите площадь круга: ")
	fmt.Scanln(&S)

	L = math.Sqrt(4 * S * math.Pi)
	fmt.Println("Длина оружности равна: ", L)

	D = math.Sqrt(4 * S / math.Pi)
	fmt.Println("Диаметр оружности равен: ", D)
}
