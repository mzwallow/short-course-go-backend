package main

import "fmt"

type Circle struct {
	radius float64
}

type CirclePtr struct {
	radius *float64
}

func main() {
	var c Circle
	fmt.Println(c)
	fmt.Printf("%T\t%p\n", c, &c)

	fmt.Println()

	cPtr := &Circle{1}
	fmt.Printf("%v\t%T\t%p\n", cPtr, cPtr, &cPtr)
	fmt.Printf("%v\t%T\t%p\n", cPtr.radius, cPtr.radius, &cPtr.radius)

	fmt.Println()

	var r float64 = 20
	a := CirclePtr{&r}
	fmt.Printf("%v\t%T\t%p\n", a, a, &a)
	fmt.Printf("%v\t%T\t%p\n", a.radius, a.radius, &a.radius)

	fmt.Println()

	var rP float64 = 10
	b := &CirclePtr{&rP}
	fmt.Printf("%v\t%T\t%p\n", b, b, &b)
	fmt.Printf("%v\t%T\t%p\n", b.radius, b.radius, &b.radius)
}
