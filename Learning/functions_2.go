package main

import "fmt"

// Vardic Functions
func sum(isAvg bool, a ...int) (tot int) {

	tot = 0
	for _, v := range a {
		tot += v
	}

	if isAvg {
		tot = tot / len(a)
	}

	return
}

func main() {

	fmt.Println("Hello, World")
	fmt.Println(sum(false, 1, 2, 3, 5, 5.0))
	fmt.Println(sum(true, 1, 2, 3, 5, 5.0))

	x := []int{1, 9, 60, 50}
	fmt.Println(sum(false, x...))
	fmt.Println(sum(true, x...)) // Spread Operator

	// Closures
	add := func(x ...int) int {
		return sum(false, x...)
	}

	fmt.Println(add(x...))

	y := 1
	inc := func() int {
		y++
		return y
	}

	fmt.Println(inc())
	fmt.Println(inc())

}
