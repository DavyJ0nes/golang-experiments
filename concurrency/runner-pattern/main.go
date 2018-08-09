package main

import (
	"log"
	"os"
	"time"

	"github.com/davyj0nes/golang-experiments/concurrency/runner-pattern/runner"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting Work")

	r := runner.New(timeout)

	// Add Three tasks to the runner
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout...")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt...")
			os.Exit(2)

		}
	}

	log.Println("Process Ended")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Task #%d - Processing\n", id)
		time.Sleep((time.Duration(id) * 500) * time.Millisecond)
		log.Printf("Task #%d - Complete\n", id)
	}
}
