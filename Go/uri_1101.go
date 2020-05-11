package main

import (
	"fmt"
	"math"
)

func main() {
	var m, n = 1, 1
	for {
		s := 0
		_, _ = fmt.Scanf("%d%d", &m, &n)
		if m <= 0 || n <= 0 {
			break
		} else {
			maxx := int(math.Max(float64(m), float64(n)))
			minn := int(math.Min(float64(m), float64(n)))
			for i := minn; i <= maxx; i++ {
				s += i
				fmt.Printf("%d ", i)
			}
			fmt.Printf("Sum=%d\n", s)
		}
	}

}
