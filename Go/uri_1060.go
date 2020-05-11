package main

import "fmt"

func main() {
	var dataArrey [6] float64
	var cont int
	for i := 0; i < len(dataArrey); i++ {
		fmt.Scan(&dataArrey[i])
		if dataArrey[i] > 0 {
			cont += 1
		}
	}
	fmt.Printf("%d valores positivos\n", cont)
}
