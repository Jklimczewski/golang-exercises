package lab1

import (
	"fmt"
	"math/rand"
	"strconv"
)

func Zad4() {
	fmt.Println("Poziom 4")
	results := make(map[string]int)
	var random int
	var attempts int
	var chosen string
	var name string
	var again string
Loop:
	for {
		random = rand.Intn(100)
		attempts = 0
		for {
			attempts++
			fmt.Print("Podaj liczbę (koniec kończy program): ")
			_, err := fmt.Scan(&chosen)
			if err == nil {
				if chosen != "koniec" {
					parsed, _ := strconv.ParseInt(chosen, 10, 0)
					if int(parsed) > random {
						fmt.Println("Za duża!")
					} else if int(parsed) < random {
						fmt.Println("Za mała")
					} else {
						fmt.Print("Brawo, podaj swoje imię: ")
						fmt.Scan(&name)
						results[name] = attempts

						fmt.Print("Jeszcze raz?: (t/n) ")
						_, err2 := fmt.Scan(&again)
						if err2 == nil {
							if again == "t" {
								break
							} else {
								for k, v := range results {
									fmt.Println(k, ": ", v)
								}
								break Loop
							}
						}
					}
				} else {
					break
				}
			}
		}
	}
}
