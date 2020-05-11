package main

import "fmt"

func main() {
	var v [5] int
	var aux = 0
	for i := 0; i < len(v); i++ {
		fmt.Scanln(&v[i])
		if v[i]%2 == 0 {
			aux += 1
		}
	}
	fmt.Printf("%d valores pares\n", aux)
}
