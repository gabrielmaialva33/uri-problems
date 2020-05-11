package main

import "fmt"

func main() {
	var n, x, y = 0, 0, 0
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d%d", &x, &y)
		if y == 0 {
			fmt.Println("divisao impossivel")
		} else {
			fmt.Printf("%.1f\n", float64(x)/float64(y))
		}
	}
}
