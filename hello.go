package main

import "fmt"

// message — параметр типа string. Функция печатает строку.
func saySomething(message string) {
	fmt.Println(message)
}

// После списка параметров указан тип результата. return возвращает значение вызывающему коду.
func greeting(name string) string {
	return "Привет, " + name + "!"
}

// a, b int — сокращённая запись для a int, b int.
func plus(a, b int) int {
	return a + b
}

// В скобках перечислены типы двух результатов.
func swap(a, b int) (int, int) {
	return b, a
}

// Именованные результаты area и perimeter доступны внутри функции как переменные.
func rectangle(width, height int) (area, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return // Возвращает текущие значения area и perimeter в этом порядке.
}

// x получает копию значения аргумента. Изменение x действует только внутри функции.
func addTwenty(x int) {
	fmt.Println("x до прибавления:", x)
	x += 20
	fmt.Println("x после прибавления:", x)
}

// ...int позволяет передать любое число аргументов типа int.
// Внутри функции numbers — срез []int, по которому можно пройти через range.
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	// Строковый литерал и результат другой функции можно передать как аргумент.
	saySomething("Привет из функции")    // Вывод: Привет из функции
	saySomething(greeting("Анна"))       // Вывод: Привет, Анна!
	fmt.Println("Сложение:", plus(3, 4)) // Вывод: Сложение: 7

	left, right := swap(1, 2)
	fmt.Println("После обмена:", left, right) // Вывод: После обмена: 2 1

	area, perimeter := rectangle(3, 4)
	fmt.Println("Площадь и периметр:", area, perimeter) // Вывод: Площадь и периметр: 12 14

	a := 8
	fmt.Println("a до вызова:", a) // Вывод: a до вызова: 8
	addTwenty(a)
	// Вывод: x до прибавления: 8
	// Вывод: x после прибавления: 28
	fmt.Println("a после вызова:", a) // Вывод: a после вызова: 8

	fmt.Println("Сумма трёх чисел:", sum(1, 2, 3)) // Вывод: Сумма трёх чисел: 6
	fmt.Println("Сумма без аргументов:", sum())    // Вывод: Сумма без аргументов: 0
	values := []int{4, 5}
	fmt.Println("Сумма элементов среза:", sum(values...)) // Вывод: Сумма элементов среза: 9
}
