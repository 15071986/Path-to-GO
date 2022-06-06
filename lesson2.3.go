package main

import "fmt"

func main() {

	var a int

	fmt.Println("Введите число: ")
	fmt.Scanln(&a)

	fmt.Println("Сотен:", +a/100)
	fmt.Println("Десятков:", +a%100/10)
	fmt.Println("Едениц:", +a%10)
}
