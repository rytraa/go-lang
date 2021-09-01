package main

import (
	"fmt"
	"os"
)

func one() {
	fmt.Println("ONE")
}

func two() {
	fmt.Println("TWO")
}

func main() {

	defer two() // Calls when the current function finishes
	one()

	f, _ := os.Open("array.go")
	defer f.Close() // Placing near open call makes it easier to understand
	// Deferred functions are run even if a runtime panic occurs

}
