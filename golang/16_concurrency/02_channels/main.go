package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Print(i, " ")
		}
		c <- 1
		c <- 2
	}()

	func() {
		for i := 99; i > 90; i-- {
			fmt.Print(i, " ")
		}
	}()
	n := <-c
	m := <-c
	o := <-c
	fmt.Println("\n", n, m, o)
}
