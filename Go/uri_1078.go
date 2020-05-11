package main

import "fmt"

func main() {
	var n = 0
	fmt.Scanln(&n)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", i, n, i*n)
	}
}
