package main

import (
	"fmt"
)

func intersection(set1, set2 []int) []int {
	var result []int

	// Create a map to store elements of set1
	set1Map := make(map[int]bool)
	for _, value := range set1 {
		set1Map[value] = true
	}

	// Iterate through set2 and add common elements to result
	for _, value := range set2 {
		if set1Map[value] {
			result = append(result, value)
		}
	}

	return result
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}
	fmt.Println(intersection(set1, set2))
}

// Пересечение двух множеств представляет собой выбор только тех элементов, которые существуют в обоих множествах.
// Для реализации пересечения множеств можно использовать цикл и временный контейнер, например, слайс.
// В цикле проходим по элементам первого множества, и для каждого элемента проверяем, есть ли он во втором множестве.
// Если да, то добавляем во временный контейнер. После завершения цикла в контейнере будет храниться пересечение двух множеств.
