var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = takeString(v, 100)
}

func takeString(s string, n int) string {
	return s[:n]
}

func main() {
	someFunc()
}

// В данном коде может возникнуть утечка памяти, поскольку строка, созданная функцией createHugeString, полностью не удаляется из памяти.

// Это можно исправить, используя вспомогательную функцию, которая принимает строку и возвращает ее подстроку.