package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for m := 1; m < 1000000; m++ {
		createFile(m)
		fmt.Println(m)
	}
}

func createFile(m int) {
	file, err := os.Create("fileName" + strconv.Itoa(m))
	if err != nil {

		defer file.Close()
	}
}
