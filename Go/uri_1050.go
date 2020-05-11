package main

import "fmt"

func main() {
	str_arrey := [9] string{"Brasilia", "Salvador", "Sao Paulo", "Rio de Janeiro", "Juiz de Fora", "Campinas", "Vitoria", "Belo Horizonte"}
	var code int
	var str_code string
	fmt.Scan(&code)
	if code == 61 {
		str_code = str_arrey[0]
	} else if code == 71 {
		str_code = str_arrey[1]
	} else if code == 11 {
		str_code = str_arrey[2]
	} else if code == 21 {
		str_code = str_arrey[3]
	} else if code == 32 {
		str_code = str_arrey[4]
	} else if code == 19 {
		str_code = str_arrey[5]
	} else if code == 27 {
		str_code = str_arrey[6]
	} else if code == 31 {
		str_code = str_arrey[7]
	} else {
		code = 0
		fmt.Println("DDD nao cadastrado")
	}
	if code != 0 {
		fmt.Println(str_code)
	}
}
