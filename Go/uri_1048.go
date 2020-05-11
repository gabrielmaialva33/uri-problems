package main

import "fmt"

func main() {
	var salary, new_salary, readjust float64
	r := [5]float64{0.04, 0.07, 0.10, 0.12, 0.15}
	var sel int
	fmt.Scan(&salary)
	if salary <= 400.00 {
		sel = 4
		readjust = salary * r[sel]

	} else if salary >= 400.01 && salary <= 800.00 {
		sel = 3
		readjust = salary * r[sel]
	} else if salary >= 800.01 && salary <= 1200.00 {
		sel = 2
		readjust = salary * r[sel]
	} else if salary >= 1200.01 && salary <= 2000.00 {
		sel = 1
		readjust = salary * r[sel]
	} else {
		sel = 0
		readjust = salary * r[sel]
	}
	new_salary = salary + readjust
	fmt.Printf("Novo salario: %.2f\nReajuste ganho: %.2f\nEm percentual: %.0f %%\n", new_salary, readjust, r[sel]*100)
}
