package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y = 0, 0
	fmt.Scanf("%d\n%d", &x, &y)
	maxx := int(math.Max(float64(x), float64(y)))
	minn := int(math.Min(float64(x), float64(y)))
	for i := minn; i < maxx; i++ {
		if (i%5 == 2 || i%5 == 3) && i > minn {
			fmt.Println(i)
		}
	}
}
