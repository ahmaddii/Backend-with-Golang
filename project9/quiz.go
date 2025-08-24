package main

import "fmt"

func main() {

	// Quiz App
	score := 0

	var (
		answer1 string
		answer2 string
		answer3 string
	)

	fmt.Println("What’s the Name of Prime Minister of Pakistan?")
	fmt.Scanln(&answer1)

	fmt.Println("What’s Pakistan’s Capital?")
	fmt.Scanln(&answer2)

	fmt.Println("Who is the Chief of Army?")
	fmt.Scanln(&answer3)

	// Checking answers separately
	if answer1 == "ShahbazSharif" {
		score += 10
	}

	if answer2 == "Islamabad" {
		score += 10
	}

	if answer3 == "AsimMunir" {
		score += 10
	}

	// Show Result
	fmt.Println("Your Total Score is:", score, "/30")
}
