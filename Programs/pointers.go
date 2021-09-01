package main

import "fmt"

func zero(x *int) {
	*x = 0
}

func main() {

	x := 4
	zero(&x)
	fmt.Println(x)

	xptr := new(int)
	zero(xptr)
	fmt.Println(*xptr)

}
