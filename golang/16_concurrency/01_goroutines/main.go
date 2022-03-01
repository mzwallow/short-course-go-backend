package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Starting Program")
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	count := 0

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Print("c:", count, " ")
			count++
		}
		wg.Done()
	}()

	go func() {
		for i := 65; i < 70; i++ {
			fmt.Print("a:", count, " ")
			count++
		}
		wg.Done()
	}()

	wg.Wait()
	// time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	fmt.Println("Count:", count)
	fmt.Println("goroutines:", runtime.NumGoroutine())
	fmt.Println("Terminating Program")
}
