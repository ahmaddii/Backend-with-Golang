package main

import "fmt"

func main() {

	// Control Flow if else

	var age int

	fmt.Println("Enter Your age :")
	fmt.Scanln(&age)

	if age >= 18 {

		fmt.Println("You are an Audlt")
	} else if age >= 13 {

		fmt.Println("You are Teenager")

	} else {
		fmt.Println("You are Minor")
	}

	// switch statements just like if else but more cleaner way

	var day int

	fmt.Println("Enter the Day Number")
	fmt.Scanln(&day)

	switch day {

	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("No Entry")

	}

	// You can also use conditions inside the switch

	number := 3

    // make sure do not write exprestion in when using swtich conditions it gives type boolean to int conversion

	switch {
	case number%2 == 0:
		fmt.Println("Even Number")
	case number%2 != 0:
		fmt.Println("Odd Number")

	}

}
