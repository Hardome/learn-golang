package main

import "fmt"

func main() {
	var count int
	fmt.Println(count) // 0

	count = 1
	fmt.Println(count) // 1

	// next объявляется, count получает новое значение в том же блоке.
	count, next := 2, 3
	fmt.Println(count, next) // 2 3
	// count, _ := 4, 5 // Ошибка компиляции. В этом блоке нет нового имени.

	{
		count := 9         // Затеняет внешний count до конца блока.
		fmt.Println(count) // 9
	}
	fmt.Println(count, next) // 2 3
}
