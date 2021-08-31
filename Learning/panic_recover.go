package main

import "fmt"

func main() {

	defer func() {
		str := recover()
		fmt.Print("The Program panicked !! cuz of ")
		fmt.Println(str)
	}()

	panic("the PANIC statement")

	fmt.Println("Hello, World") // Doesn't run this statement

	// Panic is raise a runtime error
	// Recover is used to manage it.

	// panic("PANICCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC")
	// str := recover() // Unreachable code PANIC panics and stops the execution
	// fmt.Println(str) // Solution might be to chain it with defer which runs irrespective of PANICS

}
