package main

import (
	"fmt"
	"time"
)

func channelcount(animalName string, c chan string) {

	for i := 1; i <= 5; i++ {

		c <- animalName
		time.Sleep(time.Second * 1)

	}

	close(c)

}

func main() {

	c := make(chan string, 2) // creates an buffered Channel whose capacity is 2

	//c <- "Hello World" // it waits there until some one  takes in out so it creates a deadlock

	// now you can insert beacause channel becomes buffered channel

	c <- "hello world"
	c <- "go rocks"

	// we can avoid the upper deadlock by creating seperate go routine function which works as sender

	//	go func() {

	//		c <- "hello world"
	//
	//	}()

	// now sender or reciever which is msg works parallel and syncronize

	//msg := <-c

	//fmt.Println(msg)

	//fmt.Println(<-c) // each call removes value from channel and prints it like a fifo algorithem queue
	//fmt.Println(<-c)

	msg := <-c // read value one 

	fmt.Println(msg)

	msg2 := <-c // read value 2
	fmt.Println(msg2)

	// 2nd solution to avoid this type of deadlock is to create a buffered channel in which value for channel capcith

}
