package lab3

import (
	"fmt"
)

func Helper() {
	fmt.Print("Podaj zadanie (1, 2, 3, 4, 5, 6): ")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("lab3zad1")
		Zad1()
	case 2:
		fmt.Println("lab3zad2")
		Zad2()
	case 3:
		fmt.Println("lab3zad3")
		Zad3()
	case 4:
		fmt.Println("lab3zad4")
		Zad4()
	case 5:
		fmt.Println("lab3zad5")
		Zad5()
	case 6:
		fmt.Println("lab3zad6")
		Zad6()
	}
}
