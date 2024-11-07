package main

import (
	"fmt"
	"time"
)

// loop 1 second print heavy function in background
// print 2 times
func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Heavy")
	}
}

// loop 2 second print super heavy function in background
// print 1 time
func superHeavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Super heavy")
	}
}

// Asynchronous function 
func main() {
	// call function to show in console
	go heavy()
	go superHeavy()
	fmt.Println("Fin")

	// skip the last 5th second and execute in background
	// time.Sleep(time.Second * 5)

	// to skip running without executing with select order of functions
	// select {}
}
