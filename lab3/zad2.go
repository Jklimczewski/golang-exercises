package lab3

import (
	"fmt"
)

func Zad2() {
	var mat = [2][2]int{
		{1, 2},
		{3, 1},
	}
	var mat2 = [2][2]int{
		{1, 1},
		{1, 1},
	}
	var mat3 = [2][2]int{
		{0, 3},
		{2, 1},
	}
	var sum [2][2]int
	for i, tab := range mat {
		for j := range tab {
			sum[i][j] = mat[i][j] * mat2[i][j] * mat3[i][j]
		}
	}
	fmt.Println(sum)
}
