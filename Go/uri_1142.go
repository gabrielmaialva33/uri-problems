package main

import "fmt"

func main() {
	var n = 0
	fmt.Scanf("%d", &n)
	for i := 1; i <= (n * 4); i++ {
		if !(i%4 == 0) {
			fmt.Printf("%d ", i)
		} else {
			fmt.Println("PUM")
		}
	}
}
