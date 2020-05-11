package main

import "fmt"

func main() {
	var str = ""
	for {
		fmt.Scanln(&str)
		if str == "2002" {
			fmt.Println("Acesso Permitido")
			break
		} else {
			fmt.Println("Senha Invalida")

		}
	}
}
