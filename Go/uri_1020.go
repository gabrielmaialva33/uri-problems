package main

import "fmt"

func main() {
	var idade, aux int
	fmt.Scan(&idade)
	aux = idade / 365
	fmt.Printf("%d ano(s)\n", aux)
	idade = idade % 365
	fmt.Printf("%d mes(es)\n", idade/30)
	idade = idade % 30
	fmt.Printf("%d dia(s)\n", idade)
}
