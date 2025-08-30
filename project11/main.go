package main

import (
	"errors"
	"fmt"
	"time"
)

// Varidiac Funtion

func LogRequest(details ...string) {

	timestamp := time.Now().Format("2006-01-02 15:04:05")

	fmt.Println("=========Log Request======")
	fmt.Println("Time", timestamp)

	for _, d := range details {
		fmt.Println(d)
	}

	fmt.Println("============")

}

// Step create user function using errohandling or mulitple return types

func createUser(userName, email string) (msg string, err error) {

	if userName == "" || email == "" {

		return "", errors.New("Username and Email Is Required")
	}

	if !contains(email, "@") {

		return "", errors.New("Invalid Email Address")

	}

	// Defer keyword to simulate db cleanup

	defer fmt.Println("Created User Finished")

	//final return of username or nil

	return fmt.Sprintf("User Created Successfully", userName), nil

}

// Helper Function contains used for check whether the @ is present or not in step2
// Reused function

func contains(str, substr string) bool {
	return len(str) >= len(substr) && (len(substr) == 0 || (len(str) > 0 && (str == substr || contains(str[1:], substr))))
}

// Step 4 Closure function to check the user attempts of login

func LoginCount() func() int {

	count := 0

	return func() int {

		count++

		return count
	}

}

func main() {

	msg, err := createUser("Ahmad", "malikahmad@gmail.com")

	if err != nil {

		LogRequest("Action=CreateUser", "userName=Ahmad", "Status=Failed "+ err.Error())

	} else {
		fmt.Println(msg)
		LogRequest("Action=CreateUser", "UserName=Ahmad", "Status=Success")
	}

	msg, err = createUser("", "invalidgmai.com")

	if err != nil {

		LogRequest("Action=CreateUser", "userName=unknown", "Status=failed "+ err.Error())
	} else {

		fmt.Println(msg)
	}

	// Now to check login attempst

	counter := LoginCount()

	fmt.Println("Login # :", counter())
	fmt.Println("Login # :", counter())
	fmt.Println("Login # :", counter())
}
