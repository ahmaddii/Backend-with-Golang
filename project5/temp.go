package main

import "fmt"

func main() {

	var Celcius float64
	var userInput int
	var Farhenheit float64

	fmt.Println("----Temeprature Converter----")

	fmt.Println("Enter 1 for Cel to Far or 2 for Far to Cel")

	fmt.Scanln(&userInput)

	if userInput == 1 {

		fmt.Println("Enter the Temperature in Celcius")
		fmt.Scanln(&Celcius)

		Temperature := (Celcius * 9 / 5) + 32

		fmt.Println("Temperature in Farheneit: ", float64(Temperature))

	} else if userInput == 2 {

		fmt.Println("Enter the Temperature in Farhenheit")
		fmt.Scanln(&Farhenheit)

		Temperature := (Farhenheit - 32) * 5.0 / 9.0 // to avoid integer 0 so .0 we add as float

		fmt.Println("Temperature in Celcius: ", float64(Temperature))

	} else {

		println("Plz Enter Valid Input")
	}

}
