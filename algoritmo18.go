package main

import "fmt"

func main() {
	var n int
	fmt.Print("Insira um número inteiro positivo: ")
	fmt.Scan(&n)

	primes := make([]int, 0, n)
	num := 2
	for len(primes) < n {
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, num)
		}
		num++
	}

	fmt.Printf("Os %d primeiros números primos são: %v\n", n, primes)
}
