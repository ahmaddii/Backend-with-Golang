package main

import (
	"fmt"
	"time"
)

// Simulation request to http server with varidiac function
// it helps us that every request which comes from server use within the one function

func logRequest(details ...string) {

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("Request Log")
	fmt.Println("Time", timestamp) // time is essential when getting data from server or in log requests

	for _, d := range details {
		fmt.Println(d)
	}

	fmt.Println("------------------")

}

func main() {

	logRequest("Request=Post", "path/Login", "Ahmad", "Status=200")

	logRequest("Request=Post", "path/SignUp", "Sara", "Status=201")

	logRequest("Request=Get", "path/Dashboard", "Guest", "Unauthorized")

}
