package main

import (
	"fmt"
)

func main() {
	var n int
	var x uint = 0
	var resultado uint = 1
	array := [10]int{1, 7, 71, 25, 55, 180, 33, 0, 10, 1000}
	fmt.Print("Insira um número inteiro: ")
	fmt.Scan(&n)
	for x <= 9 {
		if n == array[x] {
			resultado = 0
			break // adicionado para sair do loop
		}
		x++
	}
	if resultado == 0 {
		fmt.Print("Seu número está entre os números da lista!")
	} else {
		fmt.Print("Seu número não está entre os números da lista!")
	}
}
