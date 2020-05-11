package main

import "fmt"

func main() {
	var v [5] int
	var p, in, po, ne = 0, 0, 0, 0
	for i := 0; i < len(v); i++ {
		fmt.Scanln(&v[i])
		if v[i]%2 == 0 {
			p += 1
		} else {
			in += 1
		}
		if v[i] < 0 {
			ne += 1
		} else if v[i] == 0 {
			continue
		} else {
			po += 1
		}

	}
	fmt.Printf("%d valor(es) par(es)\n%d valor(es) impar(es)\n%d valor(es) positivo(s)\n%d valor(es) negativo(s)\n", p, in, po, ne)
}
