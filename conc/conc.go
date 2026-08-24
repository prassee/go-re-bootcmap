package conc

import (
	"fmt"
	"go-re-bootcamp/logger"
	"sync"
)

var clogger logger.Logger = logger.NewConsoleLogger("Console Logger")

func SendData(ch chan<- int) {
	for i := range 10 {
		ch <- i
	}
	defer close(ch)
}

func ReceiveData(ch <-chan int) {
	for i := range ch {
		clogger.Info(fmt.Sprintf("%d", i))
	}
}

func Worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	clogger.Info("Spawned Worker")
	defer wg.Done()
	for j := range jobs {
		results <- j * 2
	}
}

func CoOrdinate(workers int) {
	clogger.Info("preparing channels and wg")
	jobs := make(chan int)
	results := make(chan int)
	var wg sync.WaitGroup
	// spin out workers
	for range workers {
		wg.Add(1)
		go Worker(jobs, results, &wg)
	}
	// lets generate jobs & close them once all of them are sent to the channels
	go func() {
		for i := range 4 {
			jobs <- i
		}
		close(jobs)
	}()
	// collect the results

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		clogger.Info(fmt.Sprintf("computed result %d", r))
	}
}
