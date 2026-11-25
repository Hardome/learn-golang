package main

import "fmt"

type Point struct {
	X int
}

func main() {
	showAddress()
	showNilAndNew()
	showPointerParameters()
	showAddressableValues()

	// Адрес локальной переменной остаётся пригодным после возврата из функции.
	fmt.Println(*newScore(7)) // 7
}

func showAddress() {
	score := 10
	pointer := &score
	fmt.Println(*pointer) // 10

	// Присваивание копирует адрес. Оба указателя обращаются к score.
	alias := pointer
	*alias = 15
	fmt.Println(score, *pointer, pointer == alias) // 15 15 true
}

func showNilAndNew() {
	var pointer *int
	var boxed any = pointer
	// Интерфейс хранит динамический тип *int даже при нулевом значении указателя.
	fmt.Println(pointer == nil, boxed == nil) // true false
	// fmt.Println(*pointer) // Паника: разыменование nil.

	pointer = new(int)
	fmt.Println(*pointer) // 0
}

func showPointerParameters() {
	score, backup := 10, 99
	pointer := &score

	increment(pointer)
	fmt.Println(score) // 11

	// Присваивание параметру в redirect не меняет переменную pointer.
	redirect(pointer, &backup)
	fmt.Println(*pointer) // 11
	pointer = redirect(pointer, &backup)
	fmt.Println(*pointer) // 99

	// slot хранит адрес самой переменной pointer и имеет тип **int.
	slot := &pointer
	*slot = &score
	fmt.Println(*pointer) // 11
}

func increment(target *int) {
	*target = *target + 1
}

func redirect(target, replacement *int) *int {
	target = replacement
	return target
}

func showAddressableValues() {
	point := &Point{X: 3}
	// Go автоматически разыменовывает указатель при выборе поля.
	point.X = 4
	fmt.Println(point.X) // 4

	numbers := []int{1, 2}
	element := &numbers[0]
	*element = 7
	fmt.Println(numbers) // [7 2]

	for index, value := range numbers {
		// value содержит копию числа. numbers[index] обозначает сам элемент.
		copyAddress := &value
		*copyAddress += 10
		elementAddress := &numbers[index]
		*elementAddress += 1
	}
	fmt.Println(numbers) // [8 3]

	// &len(numbers) нельзя использовать: результат вызова не адресуем.
	// &scores["go"] нельзя использовать, если scores — карта.
}

func newScore(initial int) *int {
	score := initial
	return &score
}
