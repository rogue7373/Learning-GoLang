package main

import "fmt"

// This is single line comment

/* This is multi line
comment */

// This is a comment about creating a function

func main() {
	fmt.Println("Hello World")
}

// To use the below please comment out the above function with /* */

// This is a function calling an int variable

func two() {
	var myNum = 50
	fmt.Println(myNum)
}

// This function is calling a variable with a string

func three() {
	var myWord = "Hello!"
	fmt.Println(myWord)
}

// Working with multiple variables

func four() {
	var x = 5
	var y = 10
	var z = x + y
	fmt.Println(z)
}

// Creating multiple variables at the same time
func five() {
	var x, y, z = 5, 10, 15
	fmt.Println(x, y, z)
}

// Working with Strings, Ints, and Boolean

func six() {
	var myNum1 int = 90
	var myWord2 string = "Hello"
	var myBool bool = true
	fmt.Println(myNum1, myWord2, myBool)
}

// Working with an array

func seven() {
	var cars = [4]string{"Volvo", "BMW", "Ford", "Dodge"}
	fmt.Print(cars)
}

// Printing items in an array, will print the second element in an array

func eight() {
	var cars = [4]string{"Volvo", "BMW", "Ford", "Dodge"}
	fmt.Print(cars[1])
}

// Updating elements of an array

func nine() {
	var cars = [4]string{"Volvo", "BMW", "Ford", "Dodge"}
	cars[0] = "Jeep"
	fmt.Print(cars)
}

// So far this is fun
