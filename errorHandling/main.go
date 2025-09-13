package main

import (
	"errors"
	"fmt"
)

//type errors interface {

//	Error() string
//}

func divide(a, b int) (int, error) {

	if b == 0 {

		return 0, errors.New("Cant Divide with Zero")
	}

	return a / b, nil

}

func main() {

	result, err := divide(4, 4)

	if err != nil {

		fmt.Println("Error ", err)
		return
	}

	fmt.Println("Result :", result)

	// In go errors are not exception they are basically values

	// go functions often return two valus

	baseErr := errors.New("DB Error 404")
	wrappedError := fmt.Errorf("Error dut to %d", baseErr)

	fmt.Println(wrappedError)
}
