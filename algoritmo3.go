package main

import "fmt"

func main() {
	array := [4]float64{3.5, 2.8, 4.0, 8.9}
	p := 1.0
	for _, i := range array {
		p *= i
	}
	fmt.Print("O produto dos 4 termos é: ", p)
}
