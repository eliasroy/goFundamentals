package main

import "fmt"

func main() {

	//constant
	const PI = 3.14
	fmt.Println("Value of PI:", PI)

	//declaration of variable
	var name string = "John Doe"
	age := 30
	var area int

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Area:", area)

	//zero value
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("Zero value of int:", a)
	fmt.Println("Zero value of float64:", b)
	fmt.Println("Zero value of string:", c)
	fmt.Println("Zero value of bool:", d)

}
