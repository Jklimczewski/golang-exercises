package lab1

import (
	"fmt"
)

func Helper() {
	fmt.Print("Podaj zadanie (1, 2, 3, 4): ")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("lab1zad1")
		Zad1()
	case 2:
		fmt.Println("lab1zad2")
		Zad2()
	case 3:
		fmt.Println("lab1zad3")
		Zad3()
	case 4:
		fmt.Println("lab1zad4")
		Zad4()
	}
}
