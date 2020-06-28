package main

import "fmt"

func main() {
	var n, fat = 0, 1
	_, _ = fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		fat *= i
	}
	fmt.Printf("%d\n", fat)
}
