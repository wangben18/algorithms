package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker %d started job %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	jobCount := 7
	jobs := make(chan int, jobCount)
	results := make(chan int, jobCount)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= jobCount; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= jobCount; i++ {
		<-results
	}
}
