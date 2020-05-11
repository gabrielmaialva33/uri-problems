package main

import (
	"fmt"
	"os"
)

var i, g, e, n = 0, 0, 0, 0

func main() {
	n += 1
	var inter, gremio, x = 0, 0, 0
	fmt.Scanf("%d%d", &inter, &gremio)
	if inter == gremio {
		e += 1
	} else if inter > gremio {
		i += 1
	} else {
		g += 1
	}

	fmt.Println("Novo grenal (1-sim 2-nao)")
	fmt.Scanf("%d", &x)
	if x == 2 {
		fmt.Printf("%d grenais\nInter:%d\nGremio:%d\nEmpates:%d\n", n, i, g, e)
		if g > i {
			fmt.Printf("Gremio venceu mais\n")
		} else if i > g {
			fmt.Printf("Inter venceu mais\n")
		} else {
			fmt.Printf("Nao houve vencedor\n")
		}
	} else if x == 1 {
		main()
	} else {
		os.Exit(0)
	}
}
