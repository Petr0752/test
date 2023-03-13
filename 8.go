package main

import "fmt"

func setBit(num *int64, i uint, value uint) {
	if value == 1 {
		*num = *num | (1 << i)
	} else {
		*num = *num & ^(1 << i)
	}
}

func main() {
	var num int64 = 0
	setBit(&num, 0, 1)
	fmt.Printf("%d\n", num)

	setBit(&num, 0, 0)
	fmt.Printf("%d\n", num)
}

// Функция setBit принимает указатель на целое число, номер бита i и значение value (0 или 1), которое устанавливается.
// Внутри функции используются побитовые операции | (или), & (и) и ^ (исключающее или), чтобы установить или сбросить указанный бит.
