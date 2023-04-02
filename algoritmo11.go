package main

import "fmt"

func main() {
	array := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var l uint
	var c uint
	fmt.Print("Informe a linha do número que deseja consultar: ")
	fmt.Scan(&l)
	fmt.Println("Informe a coluna do número que deseja consultar: ")
	fmt.Scan(&c)
	fmt.Print("O número que se encontra na linha ", l, " e na coluna ", c, " é: ", array[l][c])
}
