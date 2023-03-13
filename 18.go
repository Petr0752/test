package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Final count:", counter.value)
}

// В начале определяется структура Counter с полем value типа int64, которое будет служить счетчиком.
// Затем определяется метод Increment, который инкрементирует значение счетчика на единицу,
// используя функцию atomic.AddInt64 из пакета sync/atomic.
// Далее определяется метод Value, который возвращает текущее значение счетчика.
// В функции main создается экземпляр структуры Counter, а затем создаются несколько горутин,
// каждая из которых вызывает метод Increment несколько раз.
// В конце функции main выводится итоговое значение счетчика, вызывая метод Value экземпляра структуры Counter.
// Использование функции atomic.AddInt64 гарантирует, что инкрементирование будет происходить в безопасной для конкурентной среды форме.
