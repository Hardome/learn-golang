package main

import (
	"fmt"
	"strings"
)

func main() {
	showCaseConversion()
	showPrefixAndSuffix()
	showTextSearch()
	showOccurrenceCount()
	showJoining()
	showReplacement()
	showSplitting()
	showTrimming()
	showSubstringPositions()
	showRunePositions()
}

func showCaseConversion() {
	fmt.Println(strings.ToUpper("Привет, Go! 42")) // ПРИВЕТ, GO! 42

	fmt.Println(strings.ToLower("Привет, Go! 42")) // привет, go! 42
	fmt.Println(strings.ToUpper("straße"))         // STRAßE — простое преобразование регистра сохраняет ß.
}

func showPrefixAndSuffix() {
	filename := "report-2026.csv"
	fmt.Println(strings.HasPrefix(filename, "report-")) // true — имя начинается с report-.

	fmt.Println(strings.HasSuffix(filename, ".csv")) // true
	fmt.Println(strings.HasSuffix(filename, ".CSV")) // false
}

func showTextSearch() {
	text := "Hello World"
	fmt.Println(strings.Contains(text, "World")) // true — подстрока есть целиком.
	fmt.Println(strings.Contains(text, "Work"))  // false — такого фрагмента нет.

	fmt.Println(strings.ContainsAny(text, "Work")) // true — в World есть W.
	fmt.Println(strings.ContainsAny(text, "xyz"))  // false — ни одной из этих букв нет.
}

func showOccurrenceCount() {
	message := "Hello World. Good bye World"
	fmt.Println(strings.Count(message, "World")) // 2
	fmt.Println(strings.Count(message, "world")) // 0 — нижний регистр в запросе.

	fmt.Println(strings.Count("aaa", "aa")) // 1

	fmt.Println(strings.Count("Пр", "")) // 3 — начало, место между буквами и конец.
}

func showJoining() {
	words := []string{"Привет", "прекрасный", "мир"}
	fmt.Println(strings.Join(words, " ")) // Привет прекрасный мир
}

func showReplacement() {
	original := "Hello World. Good bye World"
	message := strings.Replace(original, "World", "Go", 1)
	fmt.Println(message) // Hello Go. Good bye World

	fmt.Println(strings.ReplaceAll(original, "World", "Go")) // Hello Go. Good bye Go

	fmt.Println(strings.ReplaceAll("Пр", "", "-")) // -П-р- — вставка в начале и после каждой руны.
}

func showSplitting() {
	phone := "+7-987-654-32-10"
	fmt.Printf("%q\n", strings.Split(phone, "-")) // ["+7" "987" "654" "32" "10"]

	fmt.Printf("%q\n", strings.Split("7999", "-"))  // ["7999"]
	fmt.Printf("%q\n", strings.Split("a,,b,", ",")) // ["a" "" "b" ""]

	fmt.Printf("%q\n", strings.Split("П!", "")) // ["П" "!"]

	text := " \tHello  World\nGood bye World "
	fmt.Printf("%q\n", strings.Fields(text)) // ["Hello" "World" "Good" "bye" "World"]
}

func showTrimming() {
	text := "01.2300"
	fmt.Println(strings.Trim(text, "0"))      // 1.23 — удалили нули с обоих краёв.
	fmt.Println(strings.TrimLeft(text, "0"))  // 1.2300 — удалили ноль с начала.
	fmt.Println(strings.TrimRight(text, "0")) // 01.23 — удалили нули с конца.
	fmt.Println(strings.Trim("?!Go!?", "!?")) // Go — порядок рун в наборе произвольный.

	phone := "+7-987-654-32-10!"
	outsideDigit := func(symbol rune) bool {
		return symbol < '0' || symbol > '9'
	}
	fmt.Println(strings.TrimFunc(phone, outsideDigit))      // 7-987-654-32-10
	fmt.Println(strings.TrimLeftFunc(phone, outsideDigit))  // 7-987-654-32-10!
	fmt.Println(strings.TrimRightFunc(phone, outsideDigit)) // +7-987-654-32-10

	fmt.Printf("%q\n", strings.TrimSpace(" \t7 987 654 32 10\n")) // "7 987 654 32 10"

	numberLine := "+7 987 654 32 10;"
	fmt.Printf("%q\n", strings.TrimPrefix(numberLine, "+7")) // " 987 654 32 10;" — пробел сохранился.
	fmt.Printf("%q\n", strings.TrimSuffix(numberLine, ";"))  // "+7 987 654 32 10"

	fmt.Println(strings.TrimPrefix("--Go", "-")) // -Go — снят один префикс.
	fmt.Println(strings.TrimLeft("--Go", "-"))   // Go — сняты все дефисы с начала.
}

func showSubstringPositions() {
	text := "П: Go, Go!"
	offset := strings.Index(text, "Go")
	fmt.Println(offset) // 4
	if offset >= 0 {
		fmt.Println(text[offset:]) // Go, Go!
	}
	fmt.Println(strings.Index(text, "go")) // -1 — такого сочетания с малой g нет.

	fmt.Println(strings.LastIndex(text, "Go")) // 8 — начало второго Go.
}

func showRunePositions() {
	text := "П: Go, Go!"
	fmt.Println(strings.IndexAny(text, "!o")) // 5 — байтовое смещение первой o.
	fmt.Println(strings.IndexAny(text, "xy")) // -1 — ни одной из этих рун нет.

	fmt.Println(strings.LastIndexAny(text, "!o")) // 10 — байтовое смещение !.
}
