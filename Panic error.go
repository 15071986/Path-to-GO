package main

import (
	"errors"
	"fmt"
)

func div(a, b float64) (ret float64, err error) {

	defer func() {
		if e := recover(); e != nil {
			fmt.Println("err in dev2:", e)
			err = e.(error)
		}
	}()
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("err in dev:", e)
			panic(e)
		}
	}()
	if b == 0 {
		panic(errors.New("dinide by zero"))
	}
	return a / b, nil
}

func main() {
	res, err := div(4, 0)
	fmt.Println("div:", res, err)
}
