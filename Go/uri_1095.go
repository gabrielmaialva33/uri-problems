package main

import "fmt"

func main() {
	var j = 60
	for i := 1; i <= 37; i += 3 {
		fmt.Printf("I=%d J=%d\n", i, j)
		j -= 5
	}
}
