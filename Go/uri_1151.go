package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, c, N = 0, 1, 0, 0
	_, _ = fmt.Scanf("%d", &N)
	if N == 1 {
		fmt.Printf("%d\n", a)
		os.Exit(0)
	}
	fmt.Printf("%d %d ", a, b)
	for i := 2; i < N; i++ {
		c = a + b
		a = b
		b = c
		if i < N-1 {
			fmt.Printf("%d ", c)
		} else {
			fmt.Printf("%d\n", c)
		}
	}
}
