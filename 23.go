package main

import "fmt"

func main() {
	// Создаем слайс целых чисел
	slice := []int{1, 2, 3, 4, 5}

	// Индекс элемента, который нужно удалить
	i := 2

	// Удаляем элемент из слайса
	slice = append(slice[:i], slice[i+1:]...)

	// Выводим получившийся слайс
	fmt.Println(slice)
}
