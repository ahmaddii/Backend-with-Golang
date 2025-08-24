package main

import "fmt"

func main() {

	greet("Ahmad")

	fmt.Println(add(4, 3))

	divde, sub := divide(42, 10)
	fmt.Println("Divide:", divde, "Sub:", sub)

	//fmt.Println(divide(42, 4))

}

// Simple Parametrized Function

func greet(name string) {

	fmt.Println("Hello", name)

}

// Retrun type function

func add(num1 int, num2 int) int {

	return num1 + num2

}

// Go speciality ha ke wo multiple values ko bhi return kr wa skti ha

func divide(num1 int, num2 int) (int, int) {

	return num1 / num2, num1 - num2 // aik hi function mein mutiple return values go speically
}
