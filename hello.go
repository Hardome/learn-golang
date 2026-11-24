package main

import "fmt"

func main() {
	// По литералам Go выводит типы int, float64, string и bool.
	// Формат %T показывает тип, а %v — значение.
	count := 127
	price := 32.32
	label := "Привет"
	isReady := true

	fmt.Printf("count: %T = %v\n", count, count)
	fmt.Printf("price: %T = %v\n", price, price)
	fmt.Printf("label: %T = %q (bytes: %d)\n", label, label, len(label))
	fmt.Printf("isReady: %T = %v\n", isReady, isReady)

	// int8 хранит значения от -128 до 127, а uint8 — от 0 до 255.
	// byte — другое имя для uint8. rune — другое имя для int32.
	var signed int8 = -128
	var octet byte = 255
	letter := 'Ж' // Символьный литерал имеет тип rune по умолчанию.
	fmt.Printf("signed: %T = %v\n", signed, signed)
	fmt.Printf("octet: %T = %v\n", octet, octet)
	fmt.Printf("letter: %T = %c (%U)\n", letter, letter, letter)

	// float32 нужно указать явно, иначе для 0.5 будет выведен float64.
	// В комплексном числе 2 + 3i действительная часть равна 2, мнимая — 3.
	var fraction float32 = 0.5
	var smallComplex complex64 = 2 + 3i
	largeComplex := 2 + 3i // Тип по умолчанию — complex128.
	fmt.Printf("fraction: %T = %v\n", fraction, fraction)
	fmt.Printf("smallComplex: %T = %v\n", smallComplex, smallComplex)
	fmt.Printf("largeComplex: %T = %v\n", largeComplex, largeComplex)

	// При целочисленном делении дробная часть отбрасывается.
	// Преобразуем count в float64 до деления, чтобы получить 63.5.
	fmt.Printf("count / 2 = %v, float64(count) / 2 = %v\n", count/2, float64(count)/2)

	// Переменные без начального значения получают нулевое значение своего типа.
	var number int
	var text string
	var flag bool
	fmt.Printf("zero values: %v, %q, %v\n", number, text, flag)
}
