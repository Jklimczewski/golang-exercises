package lab3

import (
	"fmt"
)

func Zad1() {
	var tab [20]float64
	var tab2 [20]float64
	var sum float64
	for i := range tab {
		tab[i] = 2.0
		tab2[i] = 3.0
		sum += tab[i]
		sum += tab2[i]
	}
	fmt.Println(sum)
}
