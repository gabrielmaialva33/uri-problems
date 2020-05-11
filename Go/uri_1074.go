package main

import "fmt"

func main() {
	var n = 0
	fmt.Scanln(&n)
	v := make([]int, n)
	for i := 0; i < len(v); i++ {
		fmt.Scanln(&v[i])
	}
	for i := 0; i < len(v); i++ {
		if v[i] == 0 {
			fmt.Printf("NULL\n")
		} else if v[i]%2 == 0 {
			if v[i] > 0 {
				fmt.Printf("EVEN POSITIVE\n")
			} else {
				fmt.Printf("EVEN NEGATIVE\n")
			}
		} else {
			if v[i] > 0 {
				fmt.Printf("ODD POSITIVE\n")
			} else {
				fmt.Printf("ODD NEGATIVE\n")
			}
		}
	}
}
