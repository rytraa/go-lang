package main

import "fmt"

func main() {

	var x [5]int
	x[0] = 1
	fmt.Println(x)

	y := [5]int{10, 20, 30, 40, 50}
	fmt.Println(y)

	for i := 0; i < len(x); i++ {
		fmt.Println("i = ", x[i])
	}

	var sum float64 = 0
	for i, value := range y {
		sum += float64(value)
		fmt.Println(i, value)
	}
	fmt.Println("Avg : ", sum/float64(len(y)))

	z := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	smallest := 10000000000
	for _, val := range z {
		if val < smallest {
			smallest = val
		}
	}

	fmt.Println(smallest)

}
