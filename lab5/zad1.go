package lab5

import (
	"flag"
	"fmt"
	"math"
)

type Values struct {
	a float64
	b float64
	c float64
}
type Krotka struct {
	x1 float64
	x2 float64
}

func Calculate(p *Values) (Krotka, bool) {
	delta := ((*p).b * (*p).b) - 4*(*p).a*(*p).c
	var pierw Krotka
	if delta >= 0 {
		pdelta := math.Pow(delta, 1/2.0)
		pierw.x1 = (-1*(*p).b - pdelta) / 2 * (*p).a
		pierw.x2 = (-1*(*p).b + pdelta) / 2 * (*p).a
		return pierw, true
	} else {
		return pierw, false
	}

}
func Delta() {
	var a *float64 = flag.Float64("a", 1, "a")
	var b *float64 = flag.Float64("b", 1, "b")
	var c *float64 = flag.Float64("c", 1, "c")
	flag.Parse()
	vls := Values{*a, *b, *c}
	res, ok := Calculate(&vls)
	if ok {
		fmt.Printf("Delta dodatnia, posiada dwa miejsca zerowe x1 i x2, ktorych wartosci wynosza: %v i %v\n", res.x1, res.x2)
	} else {
		fmt.Println("Delta ujemna")
	}
}
