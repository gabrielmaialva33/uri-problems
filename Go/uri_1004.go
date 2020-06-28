package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d\n%d", &a, &b)
	fmt.Printf("PROD = %d\n", prod(a, b))
}
func prod(x, y int) int {
	return x * y
}
