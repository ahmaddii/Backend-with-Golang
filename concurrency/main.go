package main

import (
	"fmt"
	"time"
)

// Concurrency in which you run multiple task at the same time
// Structuing your program in a way so that it handle multiple tasks
// and you can switch between them efficently

// you can acheive this goroutines and channels

func printMessage(msg string) {

	for i := 1; i <= 5; i++ {

		fmt.Println(msg, i)
		time.Sleep(time.Millisecond * 500)

	}

}

func main() {

	go printMessage("Hello its go routine function") // this runs in go routine

	printMessage("This program runs in main routine") // this program runs in main routine

}
