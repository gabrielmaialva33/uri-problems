package main

import "fmt"

func main() {
	var n, in, out = 0, 0, 0
	fmt.Scanln(&n)
	v := make([]int, n)
	for _, i := range v {
		fmt.Scanln(&v[i])
		if v[i] >= 10 && v[i] <= 20 {
			in++
		} else {
			out++
		}
	}
	fmt.Printf("%d in\n%d out\n", in, out)
}
