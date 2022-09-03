package fact

import "fmt"

func Fact(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func Fact2() {
	var res float64

	for {
		var a float64
		var op string
		switch op {
		case "f":
			res = float64(Fact(int(a)))

			fmt.Scanln(&a)
			fmt.Printf("Результат выполнения операции: %.2f\n", res)

		}

	}
}
