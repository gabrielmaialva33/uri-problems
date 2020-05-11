package main

import "fmt"

func main() {
	var n, aux int
	fmt.Scan(&n)
	fmt.Printf("%d:", n / 3600)
	aux = n % 3600
	fmt.Printf("%d:", aux / 60)
	aux = n % 60
	fmt.Printf("%d\n", aux)
}
