package main

import (
	"cmp"
	"fmt"
	"slices"
)

// Немедленный вызов вычисляет начальное значение, сохраняя ports локальной.
var defaultPorts = func() []int {
	ports := []int{8080, 443}
	slices.Sort(ports)
	return ports
}()

func main() {
	// Литерал создаёт функциональное значение, которое можно вызвать через переменную.
	double := func(value int) int { return value * 2 }
	fmt.Println(double(4)) // 8

	// Аргументы после тела сразу вызывают функцию.
	fmt.Println(func(left, right int) int { return left + right }(3, 4)) // 7

	showSort()
	fmt.Println(defaultPorts) // [443 8080]
	showOperationTable()
	showNilFunction()
}

func showSort() {
	type job struct {
		name     string
		priority int
	}
	jobs := []job{{name: "backup", priority: 2}, {name: "deploy", priority: 1}}
	// Переданный компаратор задаёт порядок сортировки среза на месте.
	slices.SortFunc(jobs, func(left, right job) int {
		return cmp.Compare(left.priority, right.priority)
	})
	fmt.Println(jobs) // [{deploy 1} {backup 2}]
}

func showOperationTable() {
	type transform func(int) int
	// Карта хранит функции с одной сигнатурой под строковыми ключами.
	operations := map[string]transform{
		"square": func(value int) int { return value * value },
	}
	if square, ok := operations["square"]; ok && square != nil {
		fmt.Println(square(5)) // 25
	}
}

func showNilFunction() {
	// Функция имеет нулевое значение nil. Интерфейс при этом содержит динамический тип func().
	var callback func()
	var boxed any = callback
	fmt.Println(callback == nil, boxed == nil) // true false

	// callback() // Паника при вызове nil.
	// fmt.Println(callback == callback) // Ошибка компиляции: функции сравнимы только с nil.
	// fmt.Println(boxed == boxed) // Паника: динамический тип func() несравним.
}
