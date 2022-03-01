package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	defer fmt.Println("Hi, Mom!")
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("Recovered with message:", msg)
		}
	}()
	res, err := division(10, 0)
	if err != nil {
		panic(err)
	}
	log.Println(res)
}

func division(a, b float64) (res float64, err error) {
	if b == 0 {
		err = errors.New("division by zero")
		return
	}
	res = a / b
	return
}
