package main

import (
	"fmt"
)

func main() {

	balance := 5000

	balance = widhdraw(balance, 1000)
	fmt.Println("Available Balance :", balance)

	Order := order(1200, 1200, 0.5)
	fmt.Println("Order Total:", Order)

	CelTemp := temp(45)
	fmt.Println("Celcius temp:", CelTemp)

	loginSt := login("admin", "1234")
	fmt.Println("Login Status:", loginSt)

	fmt.Println("Add:", add(2, 3))
	fmt.Println("Sub", sub(2, 3))
	fmt.Println("Mul", mul(2, 3))
	fmt.Println("Divide", divide(2, 3))
	fmt.Println("Module", module(2, 3))

}

func widhdraw(balance, amounnt int) int {

	if amounnt > balance {

		fmt.Println("Insufficent Amount")
		return balance

	}

	return balance - amounnt

}

// for online ordering function

func order(price, quantity int, taxRate float64) float64 {

	subtotal := float64(price * quantity)
	tax := subtotal * taxRate
	return subtotal + tax

}

// temperature converter function

func temp(celcius float64) float64 {
	return (celcius * 9 / 5) + 32
}

// Login Authenticatio function with multiple return values

func login(username, password string) bool {

	return username == "admin" && password == "1234"

}

// In math calcuations

func add(a, b int) int {

	return a + b

}
func sub(a, b int) int {

	return a - b

}

func mul(a, b int) int {

	return a * b

}

func divide(a, b int) int {

	return a / b

}

func module(a, b int) int {
	return a % b
}
