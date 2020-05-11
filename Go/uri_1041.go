package main

import "fmt"

func main() {
	var x, y float64
	fmt.Scan(&x, &y)
	if x > 0 && y > 0 {
		fmt.Print("Q1\n")
	} else if x < 0 && y > 0 {
		fmt.Print("Q2\n")
	} else if x < 0 && y < 0 {
		fmt.Print("Q3\n")
	} else if x > 0 && y < 0 {
		fmt.Print("Q4\n")
	} else {
		if x != 0 && y == 0 {
			fmt.Print("Eixo X\n")
		} else if x == 0 && y != 0 {
			fmt.Print("Eixo Y\n")
		} else {
			fmt.Print("Origem\n")
		}
	}
}
