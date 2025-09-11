package main

import "fmt"

// Pointers in which we store a memory address of a variable

func main() {

	x := 10 // x ke ander 10 value ha

	Y := &x // Now Y is pointing on x what we did ke hum ne x ka memory address store krliya y ke ander

	fmt.Println("X Values is :", x)            // show x value
	fmt.Println("Memory Address of X is :", Y) // show address of x value
	fmt.Println("*Y is :", *Y)                 // value at that address

	*Y = 20 // updating x value using pointer

	fmt.Println("After Updating Using Pointer the value of X is :", *Y) // priting updated value of y

	// Similarly

	num1 := "We are the Developers of High Level"
	num2 := "We are Devops Engineers"
	num3 := "We are Bash Scripters"

	Developers := &num1 // store the num1 address ghar ka pta in developers variable
	Devops := &num2     // store num2 address in variable
	Bash := &num3       // store num3 address in bash variable

	fmt.Println(Developers)
	fmt.Println(Devops)
	fmt.Println(Bash)

	fmt.Println("Developers Values ", *Developers) // * gives values inside the address which is saved in this case the address and value both are stored in this variable
	fmt.Println("Devops Value ", *Devops)          // without * it gives the address as you know the address is stored but when used * it gives address value simple
	fmt.Println("Bash Value", *Bash)
}
