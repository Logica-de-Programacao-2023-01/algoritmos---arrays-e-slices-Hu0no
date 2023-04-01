package main

import "fmt"

func main() {
	array := [8]uint{1, 2, 3, 4, 5, 6, 7, 8}
	var x uint
	var y uint
	fmt.Println("Lista de números: ", array)
	fmt.Println("Digite o índice de um número: ")
	fmt.Scan(&x)
	fmt.Println("Agora digite o índice de outro número para trocá-lo de lugar com o primeiro: ")
	fmt.Scan(&y)
	z := array[x]
	array[x] = array[y]
	array[y] = z
	fmt.Print("A lista final com as alterações é: ", array)
}
