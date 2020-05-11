package main

import "fmt"

func main() {
	str_month := [12]string{"January", "February", "March", "April","May","June","July","August","September","October","November","December"}
	var n int
	fmt.Scan(&n)
	fmt.Printf("%s\n",str_month[n-1])
}
