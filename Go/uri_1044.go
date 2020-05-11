package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, mim, max float64
	fmt.Scan(&a, &b)
	max = math.Max(a, b)
	mim = math.Min(a, b)
	if int(max)%int(mim) == 0 {
		fmt.Print("Sao Multiplos\n")
	} else {
		fmt.Print("Nao sao Multiplos\n")
	}
}
