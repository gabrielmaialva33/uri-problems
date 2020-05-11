package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d\n%d", &a, &b)
	fmt.Printf("PROD = %d\n", PROD(a, b))
}
func PROD(x, y int) int {
	return x * y
}
