package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {

		for { // use for loop for infinite sender otherwise we get deadlock because reciever kept on waiting

			time.Sleep(time.Millisecond * 500)
			c1 <- "Hello 500 ms"

		}

	}()

	go func() {

		for {

			time.Sleep(time.Second * 2)
			c2 <- "Hello 2 Seconds value"

		}

	}()

	for { // if you want finite sender or reciever then remove this for so we receive just 2 values from multitple channels

		select { // select wait krha ha ke koi channel jo ha wo ready ho

		case msg1 := <-c1: // which ever channel msg reday it prints
			{

				fmt.Println(msg1)
			}

		case msg2 := <-c2:
			{

				fmt.Println(msg2)

			}

		}

	}

}

// Rule of thumn for sender of receier in channesl values

// are if sender stops values and receiver kept waiting
// if receiver stops and sender sends values it create deadlock
