package main

import "fmt"

func main() {
	// По литералам Go выводит типы int, float64, string и bool.
	// Формат %T показывает тип, а %v — значение.
	count := 127
	price := 32.32
	label := "Привет"
	isReady := true

	fmt.Printf("count: %T = %v\n", count, count)                         // count: int = 127
	fmt.Printf("price: %T = %v\n", price, price)                         // price: float64 = 32.32
	fmt.Printf("label: %T = %q (bytes: %d)\n", label, label, len(label)) // label: string = "Привет" (bytes: 12)
	fmt.Printf("isReady: %T = %v\n", isReady, isReady)                   // isReady: bool = true

	// int8 хранит значения от -128 до 127, а uint8 — от 0 до 255.
	// byte — другое имя для uint8. rune — другое имя для int32.
	var signed int8 = -128
	var octet byte = 255
	letter := 'Ж'                                                // Символьный литерал имеет тип rune по умолчанию.
	fmt.Printf("signed: %T = %v\n", signed, signed)              // signed: int8 = -128
	fmt.Printf("octet: %T = %v\n", octet, octet)                 // octet: uint8 = 255
	fmt.Printf("letter: %T = %c (%U)\n", letter, letter, letter) // letter: int32 = Ж (U+0416)

	// float32 нужно указать явно, иначе для 0.5 будет выведен float64.
	// В комплексном числе 2 + 3i действительная часть равна 2, мнимая — 3.
	var fraction float32 = 0.5
	var smallComplex complex64 = 2 + 3i
	largeComplex := 2 + 3i                                            // Тип по умолчанию — complex128.
	fmt.Printf("fraction: %T = %v\n", fraction, fraction)             // fraction: float32 = 0.5
	fmt.Printf("smallComplex: %T = %v\n", smallComplex, smallComplex) // smallComplex: complex64 = (2+3i)
	fmt.Printf("largeComplex: %T = %v\n", largeComplex, largeComplex) // largeComplex: complex128 = (2+3i)

	// Переменные без начального значения получают нулевое значение своего типа.
	var zero int
	var text string
	var flag bool
	fmt.Printf("zero values: %v, %q, %v\n", zero, text, flag) // zero values: 0, "", false

	// cents — отдельный тип, amount — альтернативное имя int.
	type cents int
	type amount = int
	var stored cents = 25
	var unit amount = 5
	fmt.Println("Сумма после преобразования:", int(stored)+unit) // Сумма после преобразования: 30

	// Преобразование после целочисленного деления сохраняет уже усечённый результат.
	fmt.Println(count/2, float64(count)/2, float64(count/2)) // 63 63.5 63

	number := 256
	fmt.Println(uint8(number)) // 0, младшие восемь бит числа 256 равны нулю.
	// _ = uint8(256) // Ошибка компиляции. Константа выходит за диапазон uint8.
}
