package main

import "fmt"

// Slices are a section in a array memory block

// _ _ _ _ _ _ _ _ _ _
// |________|         |
// |   x              |
// |__________________|
//         y

// y - underlying array of the slice x

// If Slice size exceeds the underlying array size
// then a new array is created and the contents are copied

func main() {

	var x []float64
	fmt.Println(len(x))

	y := make([]float64, 5)
	fmt.Println(len(y))

	arr := [5]float64{10, 20, 30, 40, 50}
	z := append(arr[0:4], 6, 10)
	fmt.Println(z)

	slice1 := []float64{20, 10, 30}
	slice2 := make([]float64, 2)
	copy(slice2, slice1)
	fmt.Println(slice2)

}
