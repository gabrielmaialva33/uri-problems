package main

import (
	"fmt"
	"math"
)

func main() {
	var n = 0
	fmt.Scanf("%d", &n)
	for i := 1; i <= (n); i++ {
		fmt.Printf("%d %d %d\n", i, int(math.Pow(float64(i), 2)), int(math.Pow(float64(i), 3)))
	}
}
