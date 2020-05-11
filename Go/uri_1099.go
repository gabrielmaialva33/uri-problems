package main

import (
	"fmt"
	"math"
)

func main() {
	var n = 0
	fmt.Scanln(&n)
	m := make([][2]int, n)
	sum := make([]int, n)

	// - input matrix
	for i := 0; i < len(m); i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&m[i][j])
		}
	}

	// - soma de impares
	for i := 0; i < len(m); i++ {
		maxx := int(math.Max(float64(m[i][0]), float64(m[i][1])))
		minn := int(math.Min(float64(m[i][0]), float64(m[i][1])))
		for j := minn; j < maxx; j++ {
			if j%2 != 0 && j != minn && j != maxx {
				sum[i] += j
			}
		}
	}

	// - exibir
	for i := range sum {
		fmt.Println(sum[i])
	}
}
