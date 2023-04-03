package main

import (
	"fmt"
)

func main() {
	var array [3][2]int
	var n int
	var linha uint = 0
	var coluna uint = 0
	for {
		if linha == 3 {
			break
		}
		fmt.Printf("Insira o elemento da linha %d e coluna %d:  ", linha, coluna)
		fmt.Scan(&n)
		array[linha][coluna] = n
		coluna++
		fmt.Printf("Insira o elemento da linha %d e coluna %d: ", linha, coluna)
		fmt.Scan(&n)
		array[linha][coluna] = n
		linha++
		coluna--
	}
	fmt.Println(array)
}
