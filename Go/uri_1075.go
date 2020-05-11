package main

import "fmt"

func main() {
	var n = 0
	fmt.Scanln(&n)
	for i := 0; i <= 10000; i++ {
		if i%n == 2 {
			fmt.Println(i)
		}
	}
}
