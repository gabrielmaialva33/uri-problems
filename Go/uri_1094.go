package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int64
	var C, R, S int64
	var totalCobaias int64
	fmt.Scanln(&N)
	m := make([][3]string, N)
	v := make([]int64, N)
	for i := 0; i < len(m); i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&m[i][j])
		}
		v[i], _ = strconv.ParseInt(m[i][0], 10, 64)
		totalCobaias += v[i]
		if m[i][1] == "C" {
			C += v[i]
		} else if m[i][1] == "R" {
			R += v[i]
		} else {
			S += v[i]
		}
	}
	fmt.Printf("Total: %d cobaias\n", totalCobaias)
	fmt.Printf("Total de coelhos: %d\n", C)
	fmt.Printf("Total de ratos: %d\n", R)
	fmt.Printf("Total de sapos: %d\n", S)
	fmt.Printf("Percentual de coelhos: %.2f %%\n", (100.0*float64(C))/float64(totalCobaias))
	fmt.Printf("Percentual de ratos: %.2f %%\n", (100.0*float64(R))/float64(totalCobaias))
	fmt.Printf("Percentual de sapos: %.2f %%\n", (100.0*float64(S))/float64(totalCobaias))
}
