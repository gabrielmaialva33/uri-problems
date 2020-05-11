package main

import "fmt"

func main() {
	var n = 0
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
