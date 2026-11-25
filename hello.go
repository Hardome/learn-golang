package main

import "fmt"

// Ключевое слово var объявляет изменяемую переменную
var someVariable string

// С типом и значением
var someValue string = "asd"

// Go выводит тип переменной из начального значения
var autoType = "string"

// Множественное объявление с типом
var a, b, c string

var (
	d string
	e int = 1
	f     = 2
)

func main() {
	// Краткое объявление := работает только внутри функций
	shortVariable := "short"

	shortVariable2, shortVariable3 := "short 2", "short 3"

	fmt.Println(shortVariable, shortVariable2, shortVariable3) // short short 2 short 3

	// Go чувствителен к регистру: name и Name — разные переменные
	name := "Go"
	Name := "Golang"
	name = "Язык Go"
	fmt.Println(name, Name) // Язык Go Golang

	// next объявляется, count получает новое значение в том же блоке
	count := 1
	count, next := 2, 3
	{
		count := 9                              // Вложенный блок объявляет собственный count
		fmt.Println("Внутренний count:", count) // Внутренний count: 9
	}
	fmt.Println("Внешний count и next:", count, next) // Внешний count и next: 2 3
}
