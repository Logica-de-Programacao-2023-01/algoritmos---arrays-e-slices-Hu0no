package main

import "fmt"

func main() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}
	var n int
	fmt.Print("O array é: ", array)
	fmt.Println("Digite um número para adicionar no primeiro e no último termo do array: ")
	fmt.Scan(&n)
	array[0] = array[0] + n
	array[6] = array[6] + n
	fmt.Print("O array resultante é: ", array)
}
