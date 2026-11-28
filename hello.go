package main

import (
	"fmt"

	"example.com/messages" // Пакет из локального каталога greetings.
)

func main() {
	// Библиотека формирует строку, а приложение выводит её на экран.
	fmt.Println(messages.Greeting("Анна")) // Привет, Анна!
}
