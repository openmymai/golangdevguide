package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside a goroutine.")
	// Declare an anonymous function and executes it as a goroutine
	go func() {
		fmt.Println("Inside a goroutine")
	}()
	fmt.Println("Outside again.")

	// Yields to the scheduler
	runtime.Gosched()
}
