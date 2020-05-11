package main

import "fmt"

func main() {
	var x, y = 1, 0
	fmt.Scanf("%d%d", &x, &y)
	for i := 1; i <= y; i++ {
		if !(i%x == 0) {
			fmt.Printf("%d ", i)
		} else {
			fmt.Printf("%d\n", i)
		}

	}
}
