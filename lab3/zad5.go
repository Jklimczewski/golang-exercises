package lab3

import (
	"fmt"
)

func Zad5() {
	wyc := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	wyc2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	sum := make([][]int, 3)
	for i := range sum {
		sum[i] = make([]int, 3)
	}

	for i, tab := range wyc {
		for j := range tab {
			var sumaTmp int = 0
			for index := range tab {
				sumaTmp += wyc[i][index] * wyc2[index][j]
			}
			sum[i][j] = sumaTmp
		}
	}
	fmt.Println(sum)
}
