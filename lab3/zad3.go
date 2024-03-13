package lab3

import (
	"fmt"
)

func Zad3() {
	wyc := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for _, v := range wyc {
		for _, el := range v {
			fmt.Printf("%d ", el)
		}
		fmt.Println()
	}

}
