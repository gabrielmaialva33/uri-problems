package main

import "fmt"

func main() {
	var n [5] float64
	for i := 0; i < 4; i++ {
		fmt.Scan(&n[i])
	}
	med := (n[0]*2.0 + n[1]*3.0 + n[2]*4.0 + n[3]*1.0) / 10.0
	fmt.Printf("Media: %.1f\n", med)
	if med >= 7.0 {
		fmt.Printf("Aluno aprovado.\n")
	} else if med >= 5.0 && med <= 6.9 {
		fmt.Printf("Aluno em exame.\n")
		fmt.Scan(&n[4])
		fmt.Printf("Nota do exame: %.1f\n", n[4])
		med = (med + n[4]) / 2.0
		if med >= 5.0 {
			fmt.Printf("Aluno aprovado.\n")
		} else {
			fmt.Printf("Aluno reprovado.\n")
		}
		fmt.Printf("Media final: %.1f\n", med)
	} else {
		fmt.Printf("Aluno reprovado.\n")
	}

}
