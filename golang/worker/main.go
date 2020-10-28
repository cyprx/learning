package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs chan int, wg *sync.WaitGroup) {
	for job := range jobs {
		fmt.Printf("worker: %v running job: %v\n", id, job)
	}
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	jobs := make(chan int)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, wg)
		wg.Add(1)
	}

	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}
