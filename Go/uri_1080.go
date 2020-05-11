package main

import (
	"fmt"
)

func main() {
	var v [100]int
	var max, p int
	for i := 0; i < len(v); i++ {
		fmt.Scanln(&v[i])
		if v[i] > max {
			max = v[i]
			p = i + 1
		}

	}
	fmt.Printf("%d\n%d\n", max, p)
}
