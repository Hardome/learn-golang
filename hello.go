package main

import "fmt"

// Без явного типа литерал 3 задаёт нетипизированную константу.
const retries = 3
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
	readMask               = 1 << iota            // Позиция 0. Значение 1.
	writeMask                                     // Позиция 1. Значение 2.
	_                                             // Позиция 2 пропущена.
	executeMask, auditMask = 1 << iota, 1 << iota // Позиция 3. Оба значения равны 8.
)

func main() {
	// Константы можно объявлять и внутри функции.
	const pageSize = 20
	fmt.Println(appName, pageSize) // Go 20

	defaultRetries := retries
	var retries8 uint8 = retries
	var retries64 int64 = retries
	const maxRetries uint8 = retries
	fmt.Printf("%T %T %T\n", defaultRetries, retries8, retries64) // int uint8 int64
	fmt.Printf("%T %d\n", maxRetries, maxRetries)                 // uint8 3
	// var invalid uint8 = 256 // Ошибка компиляции. Значение выходит за диапазон uint8.

	fmt.Println(secondsPerMinute, secondsPerHour)              // 60 3600
	fmt.Println(smallLimit, mediumLimit, largeLimit, maxLimit) // 1 1 3 3
	fmt.Println(readMask, writeMask, executeMask, auditMask)   // 1 2 8 8

	// Промежуточное значение остаётся нетипизированной константой.
	const huge = 1 << 100
	fmt.Println("После сдвига:", huge>>98) // После сдвига: 4

	// var value int64 = huge // Ошибка компиляции. huge выходит за диапазон int64.
}
