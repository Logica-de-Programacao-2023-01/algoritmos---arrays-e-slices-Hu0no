package main

import "fmt"

func main() {
	array := [10]float64{0.5, 20, 35, 3, -8, -15, 5.1, 50, 10.2, 4.8}
	slice := []float64{}
	var n int = 0
	for n < len(array) {
		if array[n] > 5 {
			slice = append(slice, array[n])
		}
		n++
	}
	fmt.Print("O slice contendo somente números maiores que 5 é: ", slice)
}
