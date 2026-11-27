package main

import (
	"fmt"

	"rsc.io/quote" // Пакет из внешнего модуля.
)

func main() {
	// Обе функции возвращают фиксированные строки.
	fmt.Println(quote.Go())  // Don't communicate by sharing memory, share memory by communicating.
	fmt.Println(quote.Opt()) // If a program is too slow, it must have a loop.
}
