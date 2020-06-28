package main

import "fmt"

func main() {
	var out []int
	var n, aux, sum = 1, 5, 0
	for {
		_, _ = fmt.Scanf("%d", &n)
		if n == 0 {
			break
		}
		sum, aux = 0, 5
		for i := n; aux > 0; i++ {
			if i%2 == 0 {
				sum += i
				aux--
			}
		}
		out = append(out, sum)
	}
	for i := 0; i < len(out); i++ {
		fmt.Printf("%d\n", out[i])
	}
}
