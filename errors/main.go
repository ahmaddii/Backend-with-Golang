package main

import (
	"errors"
	"fmt"
)

// errors

func user(email, password string) (string, error) {

	if password != "12345" {

		return "", errors.New("Plz Enter a Vaid Password")

	}

	return email, nil
}

func main() {

	result, err := user("malik@gmail.com", "12345")

	// handle if error a skta ha

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println("Welcome User: ", result)

	result2, errr := user("ali@gmail.com", "12345")

	// handle the error if comes

	if errr != nil {

		fmt.Println(errr)
		return
	}

	fmt.Println("Welcome user 2 you are here: ", result2)

}
