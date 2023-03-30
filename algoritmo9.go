package main

import "fmt"

func main() {
	array := [6]float64{1.0, 1.5, 2.0, 2.5, 3.0, 3.5}
	var n float64
	var x int = 0
	fmt.Print("Insira um número para ser adicionado em todas as posições do array: ")
	fmt.Scan(&n)
	for x < len(array) {
		array[x] = array[x] + n
		x++
	}
	fmt.Print("O array com as alterações é: ", array)
}
