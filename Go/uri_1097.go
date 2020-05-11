package main

import "fmt"

func main() {
	var x, y = 1, 7
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("I=%d J=%d\n", x, y-j)
		}
		x += 2
		y += 2
	}
}
