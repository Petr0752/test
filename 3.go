package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	result := 0
	var wg sync.WaitGroup
	wg.Add(len(numbers))

	for _, n := range numbers {
		go func(n int) {
			defer wg.Done()
			result += n * n
		}(n)
	}

	wg.Wait()
	fmt.Println("Result:", result)
}

// В этом примере каждое число из массива вычисляется в отдельной горутине, а результаты аккумулируются в переменной result.
// Когда все горутины закончат свою работу, результат можно будет вывести.

// Обратите внимание, что для синхронизации горутин используется пакет sync.
