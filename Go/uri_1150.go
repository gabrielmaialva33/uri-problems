package main

import "fmt"

func main() {
	var X, Z, count, sum = 0, 0, 0, 0
	_, _ = fmt.Scanf("%d\n%d", &X, &Z)
	for Z <= X {
		_, _ = fmt.Scanf("%d", &Z)
	}
	for i := X; sum <= Z; i++ {
		count += 1
		sum += i

	}
	fmt.Printf("%d\n", count)
}
