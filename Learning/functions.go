package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, val := range xs {
		total += val
	}
	return total / float64(len(xs))
}

func main() {

	fmt.Println("Hello, World")
	x := []float64{20, 30, 40, 50, 60}
	fmt.Println(average(x))

}
