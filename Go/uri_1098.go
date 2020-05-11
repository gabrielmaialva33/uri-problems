package main

import "fmt"

func main() {
	var x, y = 0.0, 1.0
	for i := 0; i < 11; i++ {
		for j := 0; j < 3; j++ {
			if x == 0.0 || x == 1.0 || x > 1.8 {
				fmt.Printf("I=%.0f J=%.0f\n", x, y+float64(j))
			} else {
				fmt.Printf("I=%0.1f J=%0.1f\n", x, y+float64(j))
			}
		}
		x += 0.2
		y += 0.2
	}
}
