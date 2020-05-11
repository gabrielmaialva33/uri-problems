package main

import "fmt"

func main() {
	var x, y = 0, 0
	for {
		fmt.Scanf("%d%d", &x, &y)
		if x == y {
			break
		} else {
			if x > y {
				fmt.Println("Decrescente")
			} else {
				fmt.Println("Crescente")
			}
		}
	}
}
