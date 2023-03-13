package main

import (
	"fmt"
	"reflect"
)

func detectType(i interface{}) string {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "channel"
	default:
		return "unknown"
	}
}

func main() {
	x := 123
	y := "hello"
	z := true
	w := make(chan int)

	fmt.Println(detectType(x))
	fmt.Println(detectType(y))
	fmt.Println(detectType(z))
	fmt.Println(detectType(w))
}

// Код состоит из функции determineType , которая принимает в качестве параметра значение типа interface{}.
// Далее, используется конструкция switch для проверки типа переданного значения.

// В случае, если тип переменной является целым числом (int), выводится сообщение "The value is an integer".
// Если тип переменной является строкой (string), то выводится сообщение "The value is a string".
// Также, проверяется тип переменной на значение bool и channel.

// Если же ни один из типов не совпадает, то выводится сообщение "The value is of unknown type".

// Код завершается вызовом функции determineType с разными типами данных, чтобы продемонстрировать работу программы.
