package main

import (
	"fmt"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	squaredNumbers := make(chan int)

	for _, n := range numbers {
		go func(n int) {
			squaredNumbers <- n * n
		}(n)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(<-squaredNumbers)
	}
}

// В этом примере каждое число из массива вычисляется в отдельной горутине, и результаты передаются через канал squaredNumbers.
// Когда все горутины закончат свою работу, результаты можно будет получить из канала.
