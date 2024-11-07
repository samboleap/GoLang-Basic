package main

import "fmt"


// control statements (if else and for loop)
func main() {
	// fmt.Println("Hello world")
	// // go don't have while loop
	// // if else, for, switch case, break continue

	// f := true

	// // flag equal to address of f
	// flag := &f

	// if flag == nil {
	// 	fmt.Println("Value is nil")
	// } else if *flag {
	// 	fmt.Println("True")
	// } else {
	// 	fmt.Println("False")
	// }

	arr := []string{"leap", "sambo", "middlename"}

	for i, value := range arr {
		fmt.Println(i)
		fmt.Println(value)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "leap"
	mymap["age"] = 20

	for k, v := range mymap {
		fmt.Printf("Key: %s and Value: %v", k, v)
	}

}

