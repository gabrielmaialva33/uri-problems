package main

import "fmt"

func main() {
	var N, A, S = 0, 0, 0
	_, _ = fmt.Scanf("%d%d", &A, &N)
	for N <= 0 {
		_, _ = fmt.Scanf("%d", &N)
	}
	for i := 0; i < N; i++ {
		S += A + i
	}

	fmt.Printf("%d\n", S)
}
