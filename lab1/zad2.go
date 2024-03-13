package lab1

import (
	"fmt"
)

func Zad2() {
	fmt.Print("Podaj wiek: ")
	var wiek int
	_, err := fmt.Scan(&wiek)
	if err == nil {
		dni := float32(wiek * 365)
		mercuryAge := dni / 88
		wenusAge := dni / 225
		marsAge := dni / 687
		fmt.Printf("Twoj wiek na Wenus: %v, Marsie: %v, Mercury: %v", wenusAge, marsAge, mercuryAge)

	}
}
