package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	str = strings.ToLower(str)
	charMap := make(map[rune]int)
	for _, char := range str {
		charMap[char]++
		if charMap[char] > 1 {
			return false
		}
	}
	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"
	fmt.Println(isUnique(str1)) // true
	fmt.Println(isUnique(str2)) // false
	fmt.Println(isUnique(str3)) // false
}

// Создаем функцию isUnique(str string) bool, которая будет принимать строку и возвращать true,
// Если все символы в ней уникальные, иначе false.
// Приводим строку к нижнему регистру с помощью функции strings.ToLower().
// Создаем пустую мапу charMap, которая будет использоваться для хранения количества вхождений каждого символа в строке.
// Проходим по строке символ за символом с помощью цикла for и увеличиваем значение
// Соответствующего ключа в мапе charMap[char].
// Проверяем, есть ли в мапе значение больше 1 и возвращаем false, если есть. В противном случае возвращаем true.
