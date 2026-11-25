package main

import "fmt"

func main() {
	fmt.Println(factorial(0)) // 1
	fmt.Println(factorial(5)) // 120

	fmt.Println(fibonacci(0)) // 0
	fmt.Println(fibonacci(1)) // 1
	fmt.Println(fibonacci(6)) // 8
}

// factorial вычисляет n! для небольших n. В uint64 помещается результат до 20! включительно.
func factorial(n uint) uint64 {
	// 0! и 1! равны единице. В этих случаях следующий вызов не нужен.
	if n < 2 {
		return 1
	}
	return uint64(n) * factorial(n-1)
}

// fibonacci вычисляет F(n) по двум предыдущим числам. Пример рассчитан на малые n.
func fibonacci(n uint) uint64 {
	// Первые два числа заданы. Условие также защищает от вычитания из 0 или 1.
	if n < 2 {
		return uint64(n)
	}
	// У этих вызовов есть общие подзадачи, которые вычисляются повторно.
	return fibonacci(n-1) + fibonacci(n-2)
}
