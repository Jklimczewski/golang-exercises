package lab1

import (
	"fmt"
)

func Zad1() {
	fmt.Print("Podaj wiek: ")
	var wiek int
	_, err := fmt.Scan(&wiek)
	if err == nil {
		dni := wiek * 365
		miesiecy := wiek * 12
		fmt.Println("Dni: ", dni)
		fmt.Println("Miesiecy: ", miesiecy)
	}
}
