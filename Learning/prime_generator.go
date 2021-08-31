package main

import (
	"fmt"
	"math"
)

func genPrime(n int) func() (int, bool) {

	primes := make([]bool, n+1)

	for i := range primes {
		primes[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if primes[i] {
			for j := i * i; j <= n; j += i {
				if j%i == 0 {
					primes[j] = false
				}
			}
		}
	}

	current := 0
	return func() (r int, ok bool) {

		for ; (!(primes[current])) && current < n; current++ {
		}
		r = current
		ok = current == n
		current++
		return

	}

}

func main() {

	gen := genPrime(int(math.Pow(10, 6)))
	for x, v := gen(); !v; x, v = gen() {
		fmt.Println(x)
	}

}
