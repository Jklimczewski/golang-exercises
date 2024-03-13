package lab2

import (
	"fmt"
	"os"
)

func obliczDlugosc(x int) int {
	var counter int = 0
	for x != 1 {
		counter++
		if x%2 == 0 {
			x = x / 2
		} else {
			x = x*3 + 1
		}
	}
	return counter
}

func zad3() {
	results := make(map[int]int)
	f, err := os.Create("./lab2/collatz.txt")

	for i := 1; i <= 100000; i++ {
		results[i] = obliczDlugosc(i)
		text := fmt.Sprintf("%d %d", i, results[i])

		if err == nil {
			defer f.Close()
			_, err2 := f.WriteString(text + "\n")
			if err2 != nil {
				panic(err2)
			}
		}
	}
}

func Collatz() {
	var maxDlugosc int = 0
	var maxDlugoscEl int
	var tempDlugosc int

	for i := 2000; i <= 3000; i++ {
		tempDlugosc = obliczDlugosc(i)
		if tempDlugosc > maxDlugosc {
			maxDlugosc = tempDlugosc
			maxDlugoscEl = i
		}
	}
	fmt.Printf("Najdłuższa długość dla: %d i wynosi: %d", maxDlugoscEl, maxDlugosc)
	zad3()
}
