package main

import (
	"fmt"
	"zadania/lab1"
	"zadania/lab2"
	"zadania/lab3"
	"zadania/lab5"
)

func main() {
	fmt.Print("Podaj laby (1, 2, 3, 5): ")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		lab1.Helper()
	case 2:
		lab2.Collatz()
	case 3:
		lab3.Helper()
	case 5:
		lab5.Delta()
	}
}
