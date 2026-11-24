package main

import "fmt"

// Без явного типа литерал 3 задаёт нетипизированную константу.
const defaultRetries = 3
const appName string = "Go"

const (
	secondsPerMinute int = 60
	// Константное выражение вычисляется при компиляции.
	secondsPerHour = 60 * secondsPerMinute
)

const (
	smallLimit  int = 1
	mediumLimit     // Повторяет тип int и выражение 1.
	largeLimit  int = 3
	maxLimit        // Повторяет тип int и выражение 3.
)

const (
	// iota равно индексу объявления в блоке, начиная с нуля.
	// Выражение 1 << iota повторяется с новым значением iota в каждой строке.
	readMask    = 1 << iota // 1 << 0 = 1
	writeMask               // 1 << 1 = 2
	executeMask             // 1 << 2 = 4
)

func main() {
	// Константы можно объявлять и внутри функции.
	const pageSize = 20
	fmt.Println(appName, pageSize) // Go 20

	// Нетипизированная константа подходит обоим типам, потому что 3 помещается в каждый из них.
	var retries8 uint8 = defaultRetries
	var retries64 int64 = defaultRetries
	fmt.Printf("%T %d, %T %d\n", retries8, retries8, retries64, retries64) // uint8 3, int64 3

	// const retries = retries8 // Ошибка компиляции: retries8 — переменная.
	fmt.Println(secondsPerMinute, secondsPerHour)              // 60 3600
	fmt.Println(smallLimit, mediumLimit, largeLimit, maxLimit) // 1 1 3 3
	fmt.Println(readMask, writeMask, executeMask)              // 1 2 4

	// Промежуточное значение остаётся нетипизированной константой.
	const huge = 1 << 100
	fmt.Println("После сдвига:", huge>>98) // После сдвига: 4

	const (
		_ = iota // Пропускаем 0.
		firstIndex
		secondIndex
	)
	fmt.Println("Индексы:", firstIndex, secondIndex) // Индексы: 1 2
}
