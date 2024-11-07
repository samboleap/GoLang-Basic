package main

import (
	"fmt"
	"sync"
)
func main() {
	// wait groups
	wg := &sync.WaitGroup{}
	// add, done and wait

	// add 2 lamda function in waiting group
	wg.Add(2)

	// go anonymous lamda function
	go func() {
		fmt.Println("Hello")
		wg.Done()
	}()
	go func() {
		fmt.Println("World")
		wg.Done()
	}()

	fmt.Println("Fin")
	// wait the 2 function execute
	wg.Wait()
	fmt.Println("Exit")
}
