package main

import "fmt"

// ATM Machine Simulater

func main() {

	const pin = 2287
	var balance int
	var enteredPin int

	fmt.Println("Enter Pin")
	fmt.Scanln(&enteredPin)

	if enteredPin != pin {
		println("Plz Enter a Valid Pin")
		return
	}

	if enteredPin == pin {

		println("Proceed to ATM")
	}

	fmt.Println("Plz Enter the Balance")
	fmt.Scanln(&balance)

	for {

		println("-----Allied Bank ATM--------")

		println("1. Check Balance ")
		println("2. Deposit")
		println("3. Widhdraw")
		println("4. Exit")

		var userInput int

		fmt.Println("Enter the Choice 1/4")
		fmt.Scanln(&userInput)

		if userInput == 1 {

			println("Your Current Balance is :", balance)
		}

		if userInput == 2 {

			var amount int
			fmt.Println("Enter Amount to Deposit:")
			fmt.Scanln(&amount)
			balance += amount
			fmt.Println("Amount Deposited:", amount)
			fmt.Println("Current Balance :", balance)

		}

		if userInput == 3 {

			var amount int
			fmt.Println("Enter amount to widhdraw")
			fmt.Scanln(&amount)

			if amount < balance {

				balance -= amount
				fmt.Println("Amount Widraw:", amount)
				fmt.Println("Current Balance :", balance)

			} else {
				println("--Insufficent Balance !")
			}

		}

		if userInput == 4 {

			println("Thank You for Using Allied ATM")
			break
		}

	}
}
