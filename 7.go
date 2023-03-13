package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	data := make(map[string]int)
	wg.Add(5)
	mu := &sync.Mutex{}

	for i := 0; i < 5; i++ {
		go func(i int) {
			mu.Lock()
			data[fmt.Sprintf("data_%d", i)] = i
			mu.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(data)
}

// Go не поддерживает конкурентную запись в map из нескольких горутин без использования мьютекса или другого способа синхронизации.
// Реализация конкурентной записи в map следующая:
// В данном примере мы используем мьютекс mu для синхронизации доступа к map data.
// Для этого перед записью в map вызываем mu.Lock(), а после записи - mu.Unlock().
