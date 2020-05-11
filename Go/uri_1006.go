package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scanf("%f\n%f\n%f", &a, &b, &c)
	fmt.Printf("MEDIA = %.1f\n", media(a, b, c))
}
func media(a, b, c float64) float64 {
	return (a*2.0 + b*3.0 + c*5.0) / 10.0
}
