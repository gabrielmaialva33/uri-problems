package main

import (
	"fmt"
	"sort"
)

func main() {
	var v [2] int
	var soma = 0
	fmt.Scan(&v[0], &v[1])
	s := []int{v[0], v[1]}
	sort.Ints(s)
	for i := s[0] + 1; i < s[1]; i++ {
		if i%2 != 0 {
			soma += i
		}
	}
	fmt.Println(soma)
}
