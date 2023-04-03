package main

import "fmt"

func main() {
	s := []int{}
	var t uint
	var n int
	var v uint = 1
	var soma int = 0
	fmt.Print("Insira o tamanho da lista de números: ")
	fmt.Scan(&t)
	for v <= t {
		fmt.Print("Insira um número: ")
		fmt.Scan(&n)
		s = append(s, n)
		v++
		soma += n
	}
	fmt.Println(s)
	fmt.Println("A soma da lista de números é igual a: ", soma)
}
