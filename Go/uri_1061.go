package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	// - first
	str, _ := in.ReadString('\n')
	aux := strings.Split(strings.TrimSpace(str), " ")
	d, _ := strconv.ParseInt(aux[1], 10, 64)
	str, _ = in.ReadString('\n')
	aux = strings.Split(strings.TrimSpace(str), " : ")
	h, _ := strconv.ParseInt(aux[0], 10, 64)
	m, _ := strconv.ParseInt(aux[1], 10, 64)
	s, _ := strconv.ParseInt(aux[2], 10, 64)
	tini := (d * 86400) + (h * 3600) + (m * 60) + s
	// - last
	str, _ = in.ReadString('\n')
	aux = strings.Split(strings.TrimSpace(str), " ")
	d, _ = strconv.ParseInt(aux[1], 10, 64)
	str, _ = in.ReadString('\n')
	aux = strings.Split(strings.TrimSpace(str), " : ")
	h, _ = strconv.ParseInt(aux[0], 10, 64)
	m, _ = strconv.ParseInt(aux[1], 10, 64)
	s, _ = strconv.ParseInt(aux[2], 10, 64)
	allSec := ((d * 86400) + (h * 3600) + (m * 60) + s) - tini
	fmt.Printf("%d dia(s)\n%d hora(s)\n%d minuto(s)\n%d segundo(s)\n", allSec/86400, (allSec%86400)/3600, ((allSec%86400)%3600)/60, ((allSec%86400)%3600)%60)
}

