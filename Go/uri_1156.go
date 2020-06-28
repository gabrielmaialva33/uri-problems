package main

import "fmt"

func main() {
	var S, even, odd = 1.0, 1.0, 0.0
	for i := 3; i <= 39; i++ {
		if i%2 != 0 {
			odd = float64(i)
			even += even
			S += odd / even
		}
	}
	fmt.Printf("%0.2f\n", S)
}
