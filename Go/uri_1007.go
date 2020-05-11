package main

import "fmt"

func main() {
	var num [4] int
	for i := 0; i < len(num); i++ {
		fmt.Scan(&num[i])
	}
	fmt.Printf("DIFERENCA = %d\n", dif(num[0], num[1], num[2], num[3]))
}

func dif(a, b, c, d int) int {
	return (a * b) - (c * d)
}
