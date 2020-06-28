package main

import (
	"fmt"
)

func main() {
	var n int
	var ages []int
	for {
		_, _ = fmt.Scanf("%d", &n)
		if n >= 0 {
			ages = append(ages, n)
		} else {
			break
		}
	}
	fmt.Printf("%0.2f\n", float64(sum(ages))/float64(len(ages)))
}

func sum(x []int) int {
	var sum = 0
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}
