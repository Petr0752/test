package main

import "fmt"

func swap(a *int, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func main() {
	a := 5
	b := 10

	fmt.Println("Before swap: a =", a, "b =", b)
	swap(&a, &b)
	fmt.Println("After swap: a =", a, "b =", b)
}

// В этом коде, мы используем битовую операцию XOR (^) для поменять местами значения a и b.
// Также можно использовать операцию сложения и вычитания
