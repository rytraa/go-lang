package main

import "fmt"

func sqr(x *int) {
	*x = *x * *x
}

func swap(x *int, y *int) {
	a := *x
	*x = *y
	*y = a
}

func main() {

	x := new(int)
	*x = 5
	sqr(x)

	y := new(int)
	*y = 10
	fmt.Println(*x, *y)
	swap(x, y)
	fmt.Println(*x, *y)

	v := 5
	u := 10
	fmt.Println(v, u)
	swap(&u, &v)
	fmt.Println(v, u)

}
