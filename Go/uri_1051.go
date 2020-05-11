package main

import "fmt"

func main() {
	var salary, IR float64
	ir_arrey := [3]float64{0.08, 0.18, 0.28}
	fmt.Scan(&salary)
	if salary <= 2000.00 {
		fmt.Println("Isento")
	} else if salary >= 2000.01 && salary <= 3000.00 {
		IR = (salary - 2000.00) * ir_arrey[0]
	} else if salary >= 3000.01 && salary <= 4500.00 {
		IR = (1000.00 * ir_arrey[0]) + ((salary - 3000) * ir_arrey[1]);
	} else {
		IR = (1000 * ir_arrey[0]) + (1500 * ir_arrey[1]) + ((salary - 4500) * ir_arrey[2]);
	}
	if salary > 2000.00 {
		fmt.Printf("R$ %.2f\n", IR)
	}
}
