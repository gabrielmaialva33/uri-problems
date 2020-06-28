package main

import "fmt"

func main() {
	var tempgasto, velmed int
	fmt.Scan(&tempgasto, &velmed)
	fmt.Printf("%.3f\n", (float64(tempgasto)*float64(velmed))/12.0)
}
