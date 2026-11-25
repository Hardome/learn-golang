package main

import "fmt"

type Point struct {
	X int
}

func main() {
	showAddress()
	showNilAndNew()
	showPointerToPointer()
	showPointerArray()
	showPointerCopy()
	showPointerParameters()
	showAddressableValues()

	// Адрес локальной переменной остаётся пригодным после возврата из функции.
	fmt.Println(*newScore(7)) // 7
}

func showAddress() {
	var x int = 4
	var p *int
	// &x имеет тип *int. Конкретный адрес зависит от запуска.
	p = &x
	fmt.Printf("Адрес: %p\n", p) // Например: Адрес: 0xc000010100
	fmt.Println("Значение:", *p) // Значение: 4

	// Разыменование позволяет записать новое значение в x.
	*p = 25
	fmt.Println(x) // 25

	// При коротком объявлении тип *float64 выводится из &temperature.
	temperature := 2.3
	reading := &temperature
	fmt.Println(*reading) // 2.3
}

func showNilAndNew() {
	var p *int
	fmt.Println(p == nil) // true
	// fmt.Println(*p) // Паника: разыменование nil.
	if p != nil {
		fmt.Println(*p) // Эта строка пропускается.
	}

	// Интерфейс хранит динамический тип *int даже при нулевом значении указателя.
	var boxed any = p
	fmt.Println(boxed == nil) // false

	// new(int) возвращает адрес новой переменной со значением 0.
	p = new(int)
	fmt.Println(*p) // 0
	*p = 8
	fmt.Println(*p) // 8
}

func showPointerToPointer() {
	number := 9
	p := &number
	pp := &p // Адрес переменной p, тип **int.

	fmt.Println(pp == &p, *pp == p, **pp) // true true 9

	// Через *pp можно поменять сам p. Теперь он указывает на other.
	other := 20
	*pp = &other
	fmt.Println(*p, number) // 20 9
}

func showPointerArray() {
	first, second, fourth := 1, 2, 4
	pointers := [4]*int{&first, &second, nil, &fourth}
	fmt.Println(pointers[2] == nil) // true

	for _, p := range pointers {
		// Пропускаем пустой элемент перед разыменованием.
		if p != nil {
			fmt.Println(*p) // 1, 2, 4 на отдельных строках.
		}
	}
}

func showPointerCopy() {
	score := 10
	pointer := &score

	// Присваивание копирует адрес. Оба указателя обращаются к score.
	alias := pointer
	*alias = 15
	fmt.Println(score, *pointer, pointer == alias) // 15 15 true
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
