package main

import "fmt"

// Pointers in which we store a memory address of a variable

func main() {

	x := 10 // x ke ander 10 value ha

	Y := &x // Now Y is pointing on x what we did ke hum ne x ka memory address store krliya y ke ander

	fmt.Println("X Values is :", x) // show x value
	fmt.Println("Memory Address of X is :", Y) // show address of x value 
	fmt.Println("*Y is :", *Y) // value at that address

	*Y = 20 // updating x value using pointer

	fmt.Println("After Updating Using Pointer the value of X is :", *Y)  // priting updated value of y

}
