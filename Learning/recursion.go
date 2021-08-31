package main

import "fmt"

func fact(x uint) uint {

	if x <= 1 {
		return 1
	}

	return x * fact(x-1)

}

func fibo_(x int, dp map[int]int) int {

	if v, ok := dp[x]; ok {
		return v
	}

	if x <= 1 {
		return 1
	}

	dp[x] = fibo_(x-1, dp) + fibo_(x-2, dp)
	return dp[x]

}

func fibo(x int) int {
	return fibo_(x, make(map[int]int))
}

func main() {

	fmt.Println("Hello, World")
	// fmt.Println(fact(10))
	fmt.Println(fibo(100))

}
