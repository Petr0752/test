package main

import (
	"fmt"
)

func main() {
	// Последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем карту
	stringSet := make(map[string]bool)

	// Обходим последовательность строк
	for _, s := range strings {
		// Устанавливаем значение в true для каждого уникального элемента
		stringSet[s] = true
	}

	// Выводим уникальные элементы
	for s := range stringSet {
		fmt.Println(s)
	}
}

// В Go, вы можете создать собственное множество из строк, используя карту (map),
// в которой ключи являются уникальными элементами, а значения не имеют значения.
