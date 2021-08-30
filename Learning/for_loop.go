package main

import "fmt"

func main() {

	i := 1
	for i <= 10 {
		fmt.Println("i = ", i)
		i = i + 1
	}

	fmt.Println("---------------")

	for i := 10; i > 0; i-- {
		fmt.Println("i = ", i)
	}

}
