package main

import "fmt"

func main() {
	array := [3]int{2, 8, 10}
	soma := 0
	for _, i := range array {
		soma += i
	}
	fmt.Print("A soma é: ", soma)
}
