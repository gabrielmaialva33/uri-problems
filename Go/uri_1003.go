package main

import "fmt"

func main() {
	var A, B int
	fmt.Scanf("%d\n%d", &A, &B)
	fmt.Printf("SOMA = %d\n", soma(A, B))
}
func soma(A, B int) int {
	return A + B
}
