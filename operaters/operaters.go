package main

import "fmt"

func main() {
	// Arithmetic Operators in Golang

	num1, num2 := 10, 5

	fmt.Println("Addition :", num1+num2)
	fmt.Println("Subtraction:", num1-num2)
	fmt.Println("Multiplication:", num1*num2)
	fmt.Println("Division:", num1/num2)
	fmt.Println("Modulas:", num1%num2)

	// Comparison Operators in Golng

	if num1 == num2 {

		fmt.Println(" You Numbers are not Equal : false")
	}

	if num1 != num2 {

		fmt.Println("Your Numbers are not Equal : true")
	}

	if num1 > num2 {

		fmt.Println("Yes it is Greater : true")
	}

	if num1 < num2 {

		fmt.Println("No it is not less : false")
	}

	// Logical Operators Real World Examples

	isLogin, isPasswordCorrect := true, true

	if isLogin && isPasswordCorrect { // AND Operater if both are true then it will execute to true

		fmt.Println("You are Allowed to Dashboard")
	}

	// User Permission

	isEmailVerified, isPaymentPaid := true, false

	if isEmailVerified || isPaymentPaid { // Or Operater if any one is true then it will execute

		fmt.Println("You are Allowed to Edit the Pics")
	}

	// Not Operater

	// Blocking Unauthorized Exceess

	if !isLogin { // Change true to false

		fmt.Println("Plz Login First") // to ye execute hojjae ga

	}

	//Assignment Operators

	num3 := 5

	//num3 = num3 + 5 // Normall halat mein krte hain ab ye nhi kre ge

	num3 += 5
	num3 -= 2
	num3 *= 2
	num3 /= 2
	num3 %= 2

	println(num3)

	// Real WOrld Example in Assigment Operators

	// Bank Account

	bankBalance := 1000

	// Deposit 500
	bankBalance += 500
	fmt.Println("After Deposit Current Balance :", bankBalance)

	// Widraw 200

	bankBalance -= 200
	fmt.Println("After Widthdraw Curretn Balance :", bankBalance)

	// Login Counter for the user

	Login := 0

	Login += 1

	Login += 2

	Login += 3

	fmt.Println("How Many TImes You Log In", Login)

	// Game Score Concept by Logical Operaters

	gameScore := 30

	PlayerKill := 5

	if PlayerKill >= 5 {

		gameScore += 20
		println("Final Score :", gameScore)
	}

	if PlayerKill < 5 {

		gameScore += 10
		println("Game Score :", gameScore)
	}

	// Rating Syestem for Backend of Amazon to count average reviews

	totalReveiw, numberofUsers := 0, 0

	totalReveiw += 4
	numberofUsers += 1

	totalReveiw += 3
	numberofUsers += 1

	totalReveiw += 4
	numberofUsers += 1

	totalReveiw += 2
	numberofUsers += 1

	// count the average that whats the most average reviews

	var finalReview = float64(totalReveiw) / float64(numberofUsers) // float 64 mein isliye rakha take decimal value mile

	fmt.Println("Average Reveiws:", finalReview)

}
