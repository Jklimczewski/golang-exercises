package lab3

import (
	"fmt"
)

func Zad6() {
	wyc := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	wyc2 := [][]int{
		{2, 3, 4},
		{5, 6, 7},
		{8, 9, 10},
	}

	sum := make([][]int, 3)
	for i := range sum {
		sum[i] = make([]int, 3)
	}
	sumOdwrotnie := make([][]int, 3)
	for i := range sumOdwrotnie {
		sumOdwrotnie[i] = make([]int, 3)
	}

	for i, tab := range wyc {
		for j := range tab {
			var sumaTmp int = 0
			var sumaTmp2 int = 0
			for index := range tab {
				sumaTmp += wyc[i][index] * wyc2[index][j]
				sumaTmp2 += wyc2[i][index] * wyc[index][j]
			}
			sum[i][j] = sumaTmp
			sumOdwrotnie[i][j] = sumaTmp2
		}
	}
	fmt.Println(sum)
	fmt.Printf("Suma odwrotna: %v", sumOdwrotnie)
}
