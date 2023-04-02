package main

import "fmt"

func main() {
	array := [10]uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := []uint{}
	var n int = 0
	for n < len(array) {
		if array[n]%2 == 0 {
			slice = append(slice, array[n])
		}
		n++
	}
	fmt.Print("O slice contendo somente os números pares do array é: ", slice)
}
