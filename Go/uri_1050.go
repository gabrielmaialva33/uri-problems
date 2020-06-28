package main

import "fmt"

func main() {
	strarrey := [9]string{"Brasilia", "Salvador", "Sao Paulo", "Rio de Janeiro", "Juiz de Fora", "Campinas", "Vitoria", "Belo Horizonte"}
	var code int
	var strcode string
	fmt.Scan(&code)
	if code == 61 {
		strcode = strarrey[0]
	} else if code == 71 {
		strcode = strarrey[1]
	} else if code == 11 {
		strcode = strarrey[2]
	} else if code == 21 {
		strcode = strarrey[3]
	} else if code == 32 {
		strcode = strarrey[4]
	} else if code == 19 {
		strcode = strarrey[5]
	} else if code == 27 {
		strcode = strarrey[6]
	} else if code == 31 {
		strcode = strarrey[7]
	} else {
		code = 0
		fmt.Println("DDD nao cadastrado")
	}
	if code != 0 {
		fmt.Println(str_code)
	}
}
