package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	if a < b {
		fmt.Printf("O JOGO DUROU %.0f HORA(S)\n", math.Abs(b-a))
	} else {
		fmt.Printf("O JOGO DUROU %.0f HORA(S)\n", math.Abs(24-a)+b)
	}
}
