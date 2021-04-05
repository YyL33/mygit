package main

import (
	"fmt"
	"math"
)

type Printer interface {
	Way(x,y float64) (bool)
	Scope()(x1, x2, y1, y2 float64)
	Increase()(x,y float64)
}

type Heart struct {
}

func (h Heart) Way(x, y float64) bool {
	var a float64 = x*x + y*y - 1
	return a*a*a-x*x*y*y*y <= 0
}
func (h Heart) Scope() (x1, x2, y1, y2 float64) {
	return -3, 3, -3, 3
}
func (h Heart) Increase() (x, y float64) {
	return 0.05, 0.1
}


type Sin struct {
}

func (s Sin) Way(x, y float64) bool {
	return y <= math.Sin(x)+0.1 && y >= math.Sin(x)-0.1
}
func (s Sin) Scope() (x1, x2, y1, y2 float64) {
	return -10, 10, 10, -10
}
func (s Sin) Increase() (x, y float64) {
	return 0.1, 0.2
}

func Paint(printer Printer) {
	x1, x2, y1, y2 := printer.Scope()
	xi, yi := printer.Increase()
	for y := y2; y > y1; y -= yi {
		for x := x1; x < x2; x += xi {
			if printer.Way(x, y) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	Paint(Heart{})
	//Paint(Sin{})

}
