package main

import "fmt"

func main() {
	array := [10]int{11, 52, 13, 77, 22, 702, 97, 71, 17, 107}
	var soma int = 0
	var n int = 0
	for n < len(array) {
		if n%2 == 0 {
			soma += array[n]
		}
		n++
	}
	fmt.Print("A soma dos números de índice par do array é igual a ", soma)
}
