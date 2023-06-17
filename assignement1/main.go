package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ntasks := 10

	nworkers := 5
	tasks := make(chan int,ntasks)

	wg.Add(nworkers)
	for i := 0; i < nworkers; i++ {
		go func (id int) {
			defer wg.Done()
		worker(id, tasks)

		}(i+1)
		}
	for i := 0; i < ntasks; i++ {
		tasks <- (i + 1)
	}
	close(tasks)
	startTime := time.Now()

	wg.Wait()
	fmt.Println(time.Since(startTime))
}

func worker(id int, tasks <-chan int) {
	//defer wg.Done()

	for task := range tasks {
		fmt.Printf("worker %d: starting task %d \n", id, task)
		time.Sleep(time.Second)
		fmt.Printf("worker %d :Completed task %d \n", id, task)
	}
}
