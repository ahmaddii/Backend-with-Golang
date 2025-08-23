package main

import "fmt"

func main() {

	var A = 0 // A candidate
	var B = 0 // B Candidate
	var C = 0 // C Candidate

	var userInput int
	var Voters int

	fmt.Println("Enter the Number of Voters :")
	fmt.Scanln(&Voters)

	for i := 0; i < Voters; i++ {

		fmt.Println("Wanted to Vote Press 1 for A , 2 for B and 3 For C")
		fmt.Scanln(&userInput)

		if userInput == 1 {

			A += 1
			fmt.Println("Voted Successfully")

		} else if userInput == 2 {

			B += 1
			fmt.Println("Voted Successfully")

		} else if userInput == 3 {
			C += 1
			fmt.Println("Voted Succesfully")

		} else {
			fmt.Println("Enter Valid Input !")
		}

	}

	// Now Decides Winner

	fmt.Println("Voting Summary---")
	fmt.Println("A Votes: ", A)
	fmt.Println("B Votes: ", B)
	fmt.Println("C Votes: ", C)

	// Now Decides the Winner

	if A > B && A > C {
		fmt.Println("A is the Winner")

	} else if B > A && B > C {
		fmt.Println("B is the Winner")

	} else if C > A && C > B {
		fmt.Println("C is the Winner")

	} else {
		fmt.Println("Tie !")
	}

}
