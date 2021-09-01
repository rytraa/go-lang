package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// type Android struct {
// 	Person Person
// 	Model  string
// }

// a := Android{}
// a.Person.Talk()

type Android struct {
	Person
	Model string
}

func main() {

	fmt.Println("Hello, World")
	a := Android{Person{"ffffff"}, "dkljfldkfj"}
	a.Talk()

}
