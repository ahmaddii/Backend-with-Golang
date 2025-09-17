package main

import (
	"fmt"
	"time"
)

// Learning Channels in Go
// channels are the way in golang in which we can pass values or data thorugh like a pip within in the go routine

func channelPass(thing string, c chan string) {

	for i := 1; i <= 5; i++ {

		c <- thing // c is just a string which insert value in channel comes from thing var
		time.Sleep(time.Second * 1)
	}

	close(c)

}

func main() {

	c := make(chan string) // create channel and store in variable so that we use it

	// now create go routine function

	go channelPass("Dog", c)

	go channelPass("Cat", c) // multiple go routines now we dont need wait group syncronize

	for msg := range c {

		fmt.Println(msg)
	}

}

// Go channels are like pipline in which you can pass values within in the go routine functins
