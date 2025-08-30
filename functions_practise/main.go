package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(greet("Ahmad"))

	result, err := divide(10, 0)

	if err != nil {

		fmt.Println("Error", err)
	} else {
		fmt.Println("Result", result)
	}

	userName, errr := login("Ahmad", "4324")

	if errr != nil {
		fmt.Println("Error", errr)
	} else {
		fmt.Println(userName)
	}

	userNamee, errrr := login("Ahmad", "12345")

	if errrr != nil {
		fmt.Println("Error", errr)
	} else {
		fmt.Println(userNamee)
	}

	answer := user(1, 3)
	fmt.Println(answer)

	answer2 := userr(4, 1)
	fmt.Println(answer2)

	name, age := getUserInfo(2) // beacuse two return values name or age
	fmt.Println("User Data is :", name, age)

	logEvents("Server Connected , DB connected , User logged in")

	// Defer Keyword Usage

	fmt.Println("Start")

	defer fmt.Println("File Clean Up") // Run sb se akhir mein End

	fmt.Println("End")

	HandleRequest("Ahmad")
	HandleRequest("")

	// Anoynymus Functions and Closures
    // Bina name ke function ko anonymus kehte hain
	// is ko variable mein store krke as a function used krskte hain
	greet := func(name string) string {

		return "Hello" + name
	}

	fmt.Println(greet("Ahmad"))

}

// Basic Funtion

func greet(name string) string {
	return "Hello" + name
}

// Parametrized or Multiple Return Values

func divide(num1, num2 float64) (float64, error) {

	if num2 == 0 {
		return 0, errors.New("Num2 cant be Zero") // udhar hi return krdega or show the error with new keyword
	}

	return num1 / num2, nil // nil means error ni aya

}

// You done error handling in go expicilty within the function rather than exception handling

func login(userName, Password string) (string, error) {

	if userName == "" || Password == "" {

		return "", errors.New("Username or Passwrod cant be Empty")
	}

	if Password != "12345" {
		return "", errors.New("Invalid Password")
	}

	return "Login Succesful " + userName, nil

}

// Named Retrun values in functions of Golang

func user(a, b int) int {

	return a + b

}

func userr(a, b int) (result int) {

	result = a + b
	return

}

// Example of user data fetching from database using named return functions

func getUserInfo(id int) (name, age string) {

	if id == 1 {

		name, age = "Ahmad", "18"
	} else {
		name, age = "Guest", "0"
	}

	return
}

// Variadiac functions

func logEvents(events ...string) {

	for _, e := range events {

		fmt.Println("Log :", e)
	}

}

func HandleRequest(user string) {

	fmt.Println("Handling Request for ", user)

	defer fmt.Println("Closing Request for ", user) // No matter what happend but we want to close the request when handling request in backend

	if user == "" {

		fmt.Println("Error: No user Provided")
		return
	}

	fmt.Println("Request Handle Succesfully for", user)
}

// log events = events[]string slice string ha jis mein ... string se jo values ai gi accept ho kr iss ke ander store hogi

// _ we dont need index in while looping so we ignore it
// e holds each every value everytime when loops and ... accepts mutiple strings
// range events like range givews us index or values from events slice
// then e hold that value in eveey loop and then print it in

// Defer keyword in Functions of Go
