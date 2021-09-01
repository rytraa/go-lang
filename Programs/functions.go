package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, val := range xs {
		total += val
	}
	return total / float64(len(xs))
}

func two_values() (int, int) {
	return 1, 1
}

func two_namable_return() (n int, err int) {
	n = 1
	err = 1
	return
}

func namable_return() (r int) {
	r = 1
	return
}

func main() {

	fmt.Println("Hello, World")
	x := []float64{20, 30, 40, 50, 60}
	fmt.Println(average(x))

	y, z := two_values()
	fmt.Println(y, z)

	v := namable_return()
	fmt.Println(v)

	// _ is not a valid identifier ( used to say the compiler that the value is not used )
	// _ := []float64{50, 60}
	// fmt.Println(_)

}
