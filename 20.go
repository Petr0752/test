package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun"
	fmt.Println(reverseWords(s)) // sun dog snow
}

// Разбить исходную строку на слова. Для этого можно использовать метод strings.Split(), который разбивает строку на подстроки,
// Используя разделитель.
// Обратить порядок слов. Для этого можно использовать цикл, который перебирает слова в обратном порядке и добавляет их в новую строку.
// Объединить слова в новую строку, используя пробел в качестве разделителя. Для этого можно использовать метод strings.Join().
