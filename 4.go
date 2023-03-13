package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Воркер %d совершает работу %d\n", id, j)
	}
}

func main() {
	jobs := make(chan int, 5)
	var wg sync.WaitGroup
	numWorkers := 2 // количество воркеров можно задать здесь
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// постоянная запись данных в канал основным потоком
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	// ожидание сигнала (Ctrl + C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	wg.Wait()
	fmt.Println("Все воркеры закончили работу")
}
