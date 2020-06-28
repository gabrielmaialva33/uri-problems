package main

import (
	"fmt"
)

func main() {
	var out []int
	var n, x, y, sum = 0, 0, 0, 0
	_, _ = fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		sum = 0
		_, _ = fmt.Scanf("%d%d", &x, &y)
		for j := x; y > 0; j++ {
			if j%2 != 0 {
				sum += j
				y--
			}
		}
		out = append(out, sum)
	}
	for i := 0; i < len(out); i++ {
		fmt.Printf("%d\n", out[i])
	}
}
