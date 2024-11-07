package main

import "fmt"


// interface continued
func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("vim-go")
	Anything(2.44)
	Anything("Leap")
	// empty structure
	Anything(struct{}{})


	// map(key,string), key is in string and, string can be anything 
	mymap := make(map[string]interface{})

	mymap["name"] = "L04DB4L4NC3R"
	mymap["age"] = 10
	fmt.Println(mymap)
}

