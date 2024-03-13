package lab3

import (
	"fmt"
	"math"
)

func Zad4() {
	wyc := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	wyc2 := make([][]int, 3)
	for i := range wyc2 {
		wyc2[i] = make([]int, 3)
	}

	sum := make([][]int, 3)
	for i := range sum {
		sum[i] = make([]int, 3)
	}

	for i := 2; i >= 0; i -= 1 {
		for j := 2; j >= 0; j -= 1 {
			wyc2[int(math.Abs(float64(i-2)))][int(math.Abs(float64(j-2)))] = wyc[i][j]
		}
	}

	for i, tab := range wyc {
		for i2 := range tab {
			sum[i2][i] = wyc[i2][i] + wyc2[i2][i]
		}
	}
	fmt.Println(sum)
}
