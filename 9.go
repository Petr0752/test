package main

import (
	"fmt"
)

func double(in chan int, out chan int) {
	for x := range in {
		out <- x * 2
	}
	close(out)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	input := make(chan int)
	output := make(chan int)

	go double(input, output)

	for _, number := range numbers {
		input <- number
	}
	close(input)

	for result := range output {
		fmt.Println(result)
	}
}

// В данном коде создаются 2 канала input и output.
// Функция double принимает входной канал in и выходной канал out, где каждое число x из in перемножается на 2 и записывается в out.
// В main функции создается слайс numbers с начальными числами, которые затем отправляются в input канал.
// После закрытия input канала, результаты из output канала выводятся на экран.
