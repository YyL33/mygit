package main

import "fmt"

func heart() {

	var x float64
	var y float64
	for y = 3; y > -3; y -= 0.1 {
		for x = -3; x < 3; x += 0.05 {

			var a float64 = x*x + y*y - 1
			var b float64 = a*a*a - x*x*y*y*y
			if b <= 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
func main() {
	heart()

}
