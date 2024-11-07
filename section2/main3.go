
package main

import "fmt"

// structure of interface
type Car interface {
	Drive()
	Stop()
}
// Lambo structure
type Lambo struct {
	LamboModel string
}

// create new mehtod with interface Car by assign argument in string 
func NewModel(arg string) Car {
	return &Lambo{arg}
}
// Chevy structure
type Chevy struct {
	ChevyModel string
}

// assign method stop for Lambo structure
func (l *Lambo) Stop() {
	fmt.Println("Stopping lambo")
}

// assign method drive for Lambo struture
func (l *Lambo) Drive() {
	fmt.Println("Lambo on the move")
	fmt.Println(l.LamboModel)
}

// assign method drive for Chevy structure
func (c *Chevy) Drive() {
	fmt.Println("Chevy on the move")
	fmt.Println(c.ChevyModel)
}

func main() {
	l := NewModel("Chiron")
	c := Chevy{"C369"}
	l.Drive()
	c.Drive()
}

