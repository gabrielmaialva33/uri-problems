package main

import (
	"fmt"
	"os"
)

func main() {
	var x = 1
	for {
		_, _ = fmt.Scanf("%d", &x)
		if x != 0 {
			for i := 1; i <= x; i++ {
				if i < x {
					fmt.Printf("%d ", i)
				} else {
					fmt.Printf("%d\n", i)
				}
			}
		} else {
			os.Exit(0)
		}
	}
}
