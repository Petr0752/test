package main

import (
	"fmt"
)

func quickSort(array []int, low, high int) {
	if low < high {
		pivotIndex := partition(array, low, high)
		quickSort(array, low, pivotIndex-1)
		quickSort(array, pivotIndex+1, high)
	}
}

func partition(array []int, low, high int) int {
	pivot := array[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if array[j] < pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[high] = array[high], array[i+1]
	return i + 1
}

func main() {
	array := []int{4, 65, 2, -31, 0, 99, 83, 782, 1}
	quickSort(array, 0, len(array)-1)
	fmt.Println("Sorted array:", array)
}

// Функция quickSort принимает массив целых чисел и его длину как аргументы.

// Если длина массива равна 0 или 1, то массив уже отсортирован и функция возвращает его.

// В противном случае, функция выбирает первый элемент массива как опорный элемент.

// Функция создает два массива, хранящие элементы, меньшие и большие опорного элемента.

// Функция рекурсивно вызывает саму себя для каждого из этих двух массивов.

// Наконец, функция сливает опорный элемент, массивы элементов, меньших и больших опорного элемента в один массив.

// Сортированный массив возвращается.
