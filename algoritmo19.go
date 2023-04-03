package main

import "fmt"

func main() {
	var n int
	var n2 int
	var t int
	var i int = 0
	fmt.Print("Insira o tamanho do array: ")
	fmt.Scan(&t)
	array := []int{}
	array2 := []int{}
	for i < t {
		fmt.Print("Insira os números do primeiro array: ")
		fmt.Scan(&n)
		array = append(array, n)
		i++
	}
	i = 0
	for i < t {
		fmt.Print("Insira os números do segundo array: ")
		fmt.Scan(&n2)
		array2 = append(array2, n2)
		i++
	}
	array3 := make([]int, 0, 2*t)
	array3 = append(array3, array...)
	array3 = append(array3, array2...)
	fmt.Print("O array que contém a soma dos dois arrays é: ", array3)
}
