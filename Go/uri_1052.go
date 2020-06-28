package main

import "fmt"

func main() {
	strmonth := [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	var n int
	fmt.Scan(&n)
	fmt.Printf("%s\n", strmonth[n-1])
}
