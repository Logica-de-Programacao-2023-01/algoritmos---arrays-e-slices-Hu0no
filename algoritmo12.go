package main

import "fmt"

func main() {
	array := [5]int{6, 10, 15, 18, 59}
	slice := []int{}
	var x int = 0
	for x < len(array) {
		if array[x]%3 == 0 {
			slice = append(slice, array[x])
		}
		x++
	}
	fmt.Print("O slice que contém somente os múltilos de 3 é: ", slice)
}
