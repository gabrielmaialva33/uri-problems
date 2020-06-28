package main

import "fmt"

func main() {
	var S = 1.0
	for i := 2; i <= 100; i++ {
		S += 1.0 / float64(i)
	}
	fmt.Printf("%0.2f\n", S)
}
