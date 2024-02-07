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

	// ints
	var age1 int = 30
	var age2 = 40
	age3 := 50
	var age4 int // This will be 0

	fmt.Println(age1, age2, age3, age4)

	// bits & memory

	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64
	// byte is an alias for uint8
	// rune is an alias for int32
	// int is 32 or 64 bits depending on the system
	// uint is 32 or 64 bits depending on the system

}
