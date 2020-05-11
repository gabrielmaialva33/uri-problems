package main

import "fmt"

func main() {
	var temp_gasto, vel_med int
	fmt.Scan(&temp_gasto, &vel_med)
	fmt.Printf("%.3f\n", (float64(temp_gasto)*float64(vel_med))/12.0)
}
