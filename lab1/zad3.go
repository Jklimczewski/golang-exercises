package lab1

import (
	"fmt"
	"math"
)

func Zad3() {
	fmt.Print("Podaj parametry funkcji: ")
	var a float64
	var b float64
	var c float64
	_, err := fmt.Scan(&a, &b, &c)
	if err == nil {
		fmt.Printf("a: %f, b: %f, c: %f", a, b, c)
		fmt.Println()
		delta := (b * b) - 4*(a*c)
		if delta > 0 {
			pdelta := math.Pow(delta, 1/2.0)
			x1 := (-b - pdelta) / 2 * a
			x2 := (-b + pdelta) / 2 * a
			fmt.Printf("Delta dodatnia, posiada dwa miejsca zerowe x1 i x2, ktorych wartosci wynosza: %f  i  %f", x1, x2)
		} else if delta == 0 {
			x3 := -b / 2 * a
			fmt.Printf("Delta wynosi 0, posiada więc ona jedno miejsce zerowe, ktorej wartosc to %f", x3)
		} else {
			fmt.Println("Brak pierwiastków")
		}
	}
}
