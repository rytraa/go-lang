package main

import "fmt"

// Integers

// uint8 (byte) uint16 uint32 uint64
// int8 int16 int32 (rune) int64

// Floating Point Numbers

// float32 float64
// complex64 complex128

// String

// "" (single line strings) or `` (multiline strings)

// Boolean (1 byte)

// true - 1 | false - 0

func main() {

	fmt.Println(1+1, 5*3, 15/3, 15%6)
	fmt.Println(1.0 + 2.0)
	fmt.Println("Single line \\n is used for newline \\t for tab: \nnewline \t tabbed!")
	fmt.Println(`Multi line
newline     tabbed!`)

	fmt.Println(len("Hello, world"))
	fmt.Println("helloworld"[1])
	fmt.Println("Hello," + "world")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
