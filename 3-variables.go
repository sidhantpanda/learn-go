package main

import "fmt"

func main() {
	// Different ways to declare variables
	var something1 string = "IronMan"
	var something2 = "Superman"
	var something3 string // This will be an empty string

	fmt.Println(something1, something2, something3)

	// Different ways to declare multiple variables
	var (
		name1 = "IronMan"
		name2 = "Superman"
		name3 = "Batman"
	)

	fmt.Println(name1, name2, name3)

	// Short variable declaration
	// This is the most concise way to declare variables
	// It is only available inside a function
	// It is not available outside a function

	name4 := "IronMan"
	name5 := "Superman"
	name6 := "Batman"

	fmt.Println(name4, name5, name6)
}
