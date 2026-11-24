package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 7, 2
	// Результат сравнения — логическое значение.
	// Сравнение ==
	fmt.Println("7 == 2 =", a == b) // 7 == 2 = false
	// Не равно
	fmt.Println("7 != 2 =", a != b) // 7 != 2 = true
	// Меньше
	fmt.Println("7 < 2 =", a < b) // 7 < 2 = false
	// Меньше либо равно
	fmt.Println("7 <= 7 =", a <= 7) // 7 <= 7 = true
	// Больше
	fmt.Println("7 > 2 =", a > b) // 7 > 2 = true
	// Больше либо равно
	fmt.Println("2 >= 7 =", b >= a) // 2 >= 7 = false

	age := 20
	hasTicket := false
	isAdult := age >= 18
	canEnter := isAdult && hasTicket                            // Для входа нужны оба условия.
	fmt.Println("isAdult && hasTicket =", canEnter)             // isAdult && hasTicket = false
	fmt.Println("isAdult || hasTicket =", isAdult || hasTicket) // isAdult || hasTicket = true
	fmt.Println("!hasTicket =", !hasTicket)                     // !hasTicket = true

	// При ложной левой части && правая часть пропускается. Деления на ноль здесь нет.
	divisor := 0
	isBigEnough := divisor != 0 && a/divisor > 2
	fmt.Println("Безопасная проверка =", isBigEnough) // Безопасная проверка = false

	// if принимает любое выражение типа bool, в том числе готовую переменную.
	var message string // ""

	if canEnter {
		message = "Вход разрешён"
	} else {
		message = "Вход закрыт"
	}

	fmt.Println(message) // Вход закрыт

	// Присваивание nil-указателя интерфейсу сохраняет динамический тип *int.
	var pointer *int
	var boxed any = pointer
	fmt.Println("Указатель и интерфейс равны nil:", pointer == nil, boxed == nil) // Указатель и интерфейс равны nil: true false

	nan := math.NaN()
	fmt.Println("NaN равен себе:", nan == nan) // NaN равен себе: false
}
