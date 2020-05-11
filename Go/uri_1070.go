package main

import "fmt"

func main() {
	var n = 0
	fmt.Scanln(&n)
	for i := n; i < n+12; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
