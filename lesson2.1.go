package main

import "fmt"

func main() {
	var lenght int
	var width int

	fmt.Println("Введите длину:")
	fmt.Scanln(&lenght)

	fmt.Println("Введитеи ширину:")
	fmt.Scanln(&width)

	fmt.Println("Площадь прямоугольника равна:")
	fmt.Println(lenght * width)

}
