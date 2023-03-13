package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 20) // инициализация переменной a значением 2^20
	b := big.NewInt(3 << 20) // инициализация переменной b значением 3*2^20

	// перемножение
	c := new(big.Int)
	c.Mul(a, b)
	fmt.Printf("%v * %v = %v\n", a, b, c)

	// деление
	d := new(big.Int)
	d.Div(b, a)
	fmt.Printf("%v / %v = %v\n", b, a, d)

	// сложение
	e := new(big.Int)
	e.Add(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, e)

	// вычитание
	f := new(big.Int)
	f.Sub(b, a)
	fmt.Printf("%v - %v = %v\n", b, a, f)
}
