package main

import "fmt"

func main() {

	i := 10
	if i > 10 {
		fmt.Println("Greater than 10")
	} else {
		fmt.Println("Less or equal to 10")
	}

	fmt.Println("------------------")

	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}

}
