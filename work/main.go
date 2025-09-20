package main

import (
	"fmt"
	"time"
)

func worker(Id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {

		fmt.Println("Worker", Id, "started job", j) // job start hoi
		time.Sleep(time.Second * 2)                 // 2 second bad finishd hogai
		fmt.Println("Worker", Id, "finished job", j)
		results <- j * 2 // job finished ho ke results channel mein gai

	}
}

func main() {

	numJobs := 5 // USER NE 5 chezon ke request bheji ha

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// create three workers which perform 5 jobs

	for w := 1; w <= 3; w++ {

		go worker(w, jobs, results)

	}

	// now send 5 jobs into the job channel so that worker know that

	for j := 1; j <= numJobs; j++ {

		jobs <- j

	}

	close(jobs) // no more jobs are coming in this channel

	// receive results after completing job

	for r := 1; r <= numJobs; r++ {

		<-results

	}

}

//Why Worker Pools Are Useful in Backends

//✅ Efficiency → multiple requests/jobs handled at once

//✅ Load control → you decide how many workers (so the server doesn’t get overloaded)

//✅ Scalability → can easily increase workers if traffic grows

//✅ Fairness → jobs are distributed evenly among workers
