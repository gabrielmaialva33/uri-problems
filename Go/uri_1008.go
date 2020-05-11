package main

import "fmt"

func main() {
	var num, h int
	var value float64
	_, _ = fmt.Scan(&num, &h, &value)
	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", num, float64(h)*value)
}
