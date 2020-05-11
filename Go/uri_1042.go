package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	v := []int{a, b, c}
	sort.Ints(v)
	for i := 0; i < len(v); i++ {
		fmt.Println(v[i])
	}
	fmt.Printf("\n%d\n%d\n%d\n", a, b, c)
}
