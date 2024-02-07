package main

import "fmt"

func main() {
	age := 30
	name := "Sidhant"

	fmt.Print("Hello, ")
	fmt.Print("World! \n")

	fmt.Println("new line ")

	fmt.Print("my age is ", age, " and my name is ", name, "\n")

	// Formatted strings

	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)

	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("You scored %f points \n", 153.351)
	fmt.Printf("You scored %0.1f points \n", 153.351) // Round off to 1 decimal place

	var str = fmt.Sprintf("my age is %v and my name is %v", age, name)
	fmt.Println(str)
}
