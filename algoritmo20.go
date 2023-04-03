package main

import "fmt"

func main() {
	var n int
	array := []int{}
	for {
		fmt.Print("Insira os números do array: ")
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		array = append(array, n)
	}

	crescente := true
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			crescente = false
			break
		}
	}
	if crescente == true {
		fmt.Print("O array está em ordem crescente!")
	} else {
		fmt.Print("O array não está em ordem crescente!")
	}
}
