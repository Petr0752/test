package main

import (
	"fmt"
	"time"
)

func writeToChannel(ch chan int, seconds int) {
	for i := 0; i < seconds; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func readFromChannel(ch chan int) {
	for value := range ch {
		fmt.Println("Read value:", value)
	}
	fmt.Println("Channel closed.")
}

func main() {
	ch := make(chan int)
	seconds := 5

	go writeToChannel(ch, seconds)
	go readFromChannel(ch)

	time.Sleep(time.Duration(seconds+1) * time.Second)
}

// Эта программа создает канал, затем запускает два горутины: одна для записи в канал, а другая для чтения из канала.
// В горутине writeToChannel за seconds секунд отправляются значения в канал, а в горутине readFromChannel эти значения читаются и выводятся в stdout.

// // В конечном итоге в главной функции main вызывается time.Sleep на время seconds+1 секунды, чтобы дать возможность горутинам завершиться.
