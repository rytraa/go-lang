package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scanf("%d", &n)

	for i := 1; i < n; i++ {

		var did bool = false

		if i%3 == 0 {
			print("Fizz")
			did = true
		}

		if i%5 == 0 {
			did = true
			print("Buzz")
		}

		if !did {
			print(i)
		}

		print("\n")

	}

}
