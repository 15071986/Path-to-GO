package sort

import (
	"fmt"
)

func Sort() {
	ar := []int{4, 2, 3, 1}
	InsertionSort(ar)
	fmt.Println(ar)
}

func InsertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
}
