package main

import (
	"fmt"
	"math"
)

func main() {
	var v [4] int
	/* V[0] - Hora Inicio
	   V[1] - Minuto Inicio
	   V[2] - Hora Termino
	   V[3] - Minuto Termino
	*/
	var dh, dm int
	for i := 0; i < len(v); i++ {
		fmt.Scan(&v[i])
	}
	ini := v[1] + (v[0] * 60)
	fim := v[3] + (v[2] * 60)
	if ini >= fim {
		dh = ((1440 - ini) + fim) / 60
	} else {
		dh = (fim - ini) / 60
	}
	if v[1] > v[3] {
		dm = (60 - v[1]) + v[3]
	} else {
		dm = int(math.Abs(float64(v[3] - v[1])))
	}
	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", dh, dm)
}
