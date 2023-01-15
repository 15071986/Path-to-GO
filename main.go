package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for m := 1; m < 5; m++ {
		CreateFile(m)
		fmt.Println(m)
	}
}

func CreateFile(m int) {
	file, err := os.Create("Oleg" + strconv.Itoa(m))
	if err != nil {

		defer file.Close()
	}
}
