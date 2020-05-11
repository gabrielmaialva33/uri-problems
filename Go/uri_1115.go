package main

import "fmt"

func main() {
	var x, y = 0, 0
	for {
		fmt.Scanln(&x, &y)
		if x > 0 && y > 0 {
			fmt.Println("primeiro")
		} else if x < 0 && y > 0 {
			fmt.Println("segundo")
		} else if x < 0 && y < 0 {
			fmt.Println("terceiro")
		} else if x > 0 && y < 0 {
			fmt.Println("quarto")
		} else {
			break
		}
	}
}
