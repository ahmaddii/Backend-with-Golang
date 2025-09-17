package main

import (
	"fmt"
	"sync"
	"time"
)

// Concurrency in which you run multiple task at the same time
// Structuing your program in a way so that it handle multiple tasks
// and you can switch between them efficently

// you can acheive this goroutines and channels

//func printMessage(msg string) {

//	for i := 1; i <= 5; i++ {

//		fmt.Println(msg, i)
//		time.Sleep(time.Millisecond * 500)

//	}

//}

func infiniteLoop(animalName string) {

	for i := 0; i <= 5; i++ {

		fmt.Println(i, animalName)
		time.Sleep(time.Second * 1)
	}

}

func main() {

	//	go printMessage("Hello its go routine function") // this runs in go routine

	//	printMessage("This program runs in main routine") // this program runs in main routine

	//go infiniteLoop("Cat") // runs in background
	//go infiniteLoop("Dog") // runs in main routine // non blocking nature of goroutines

	// Asynroncaliy cause empty output beacause it mains die immealidely

	// Now we use wait group to syncronize our program

	var wg sync.WaitGroup // initialze wait group counter

	wg.Add(2) // like a counter in which we know 2 go rountines functions are coming

	go func() {

		infiniteLoop("Dog")
		wg.Done() // tells the wait group 1 go routine is done
	}()

	// 2nd go routine is done

	go func() {

		infiniteLoop("Cat")
		wg.Done() // 2nd bhi done
	}()

	wg.Wait() // wait karo main until mere go routine wale functions chal na jaye

}

//wg.Add(1) → counter = 1 (we’re expecting 1 goroutine).

//The goroutine runs count("dog") → prints 5 times.

//After finishing, it calls wg.Done() → counter = 0.

///wg.Wait() sees counter = 0 → main continues → program exits safely.
