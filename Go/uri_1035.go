package main

import "fmt"

func main() {
	var v [4] int
	for i := 0; i < len(v); i++ {
		fmt.Scan(&v[i])
	}
	if v[1] > v[2] && v[3] > v[0] && (v[2]+v[3]) > (v[0]+v[1]) && v[2] >= 0 && v[3] >= 0 && v[0]%2 == 0 {
		fmt.Printf("Valores aceitos\n")
	} else {
		fmt.Printf("Valores nao aceitos\n")

	}
}
