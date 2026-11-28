package main

import "fmt"

func main() {
	total := 10
	total = 12
	fmt.Println(total) // 12

	// Правая часть читает оба прежних значения до записи.
	left, right := "лево", "право"
	left, right = right, left
	fmt.Println(left, right) // право лево

	// Элемент массива выбирается по прежнему индексу 0.
	items := [2]int{10, 20}
	index := 0
	index, items[index] = 1, 99
	fmt.Println(index, items) // 1 [99 20]

	total += 3
	fmt.Println(total) // 15

	var flags uint8 = 0b0101
	flags |= 0b0010
	fmt.Printf("%04b\n", flags) // 0111

	items = [2]int{10, 20}
	calls := 0
	nextIndex := func() int {
		calls++
		return 0
	}
	items[nextIndex()] += 5
	fmt.Println(items, calls) // [15 20] 1
}
