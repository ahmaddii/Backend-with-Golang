package main

import "fmt"

func main() {

	// Grading Syestem of Szabist

	var Marks int

	fmt.Println("Enter the Marks")
	fmt.Scanln(&Marks)

	if Marks >= 90 {

		fmt.Println("A Grade")

	} else if Marks >= 80 && Marks < 90 {

		fmt.Println("B Grade")

	} else if Marks >= 70 && Marks < 80 {

		fmt.Println("C Grade")

	} else if Marks >= 60 && Marks < 70 {

		fmt.Println("D Grade")
	} else {

		fmt.Println("Fail")
	}

}
