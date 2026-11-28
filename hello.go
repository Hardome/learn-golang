package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	showStringLiteralsAndIndexing()
	showRunes()
	showStringConversions()
	showIteration()
	showLengths()
	showStringReplacement()
	showSubstrings()
}

func showStringLiteralsAndIndexing() {
	// В обратных кавычках \n остаётся последовательностью из двух байтов.
	fmt.Printf("%q\n", "Go\n") // "Go\n"
	fmt.Printf("%q\n", `Go\n`) // "Go\\n"

	ascii := "Hello"
	fmt.Println(len(ascii), ascii[0]) // 5 72

	text := "Привет"
	fmt.Println(len(text), text[0]) // 12 208 — шесть букв по два байта, первый байт П равен 208.
}

func showRunes() {
	var letter rune = 'П'       // Кодовая точка U+041F.
	fmt.Println(letter)         // 1055
	fmt.Printf("%c\n", letter)  // П
	fmt.Println(string(letter)) // П

	// Байты D0 9F кодируют П, а число 0xD09F указывает на другую кодовую точку.
	var bytePairAsRune rune = 0xD09F
	fmt.Printf("%c %U\n", bytePairAsRune, bytePairAsRune) // 킟 U+D09F
}

func showStringConversions() {
	text := "Привет"
	bytes := []byte(text)
	fmt.Println(bytes)      // [208 159 209 128 208 184 208 178 208 181 209 130]
	fmt.Println(len(bytes)) // 12

	letters := []rune(text)
	fmt.Println(letters)           // [1055 1088 1080 1074 1077 1090]
	fmt.Println(len(letters))      // 6
	fmt.Printf("%c\n", letters[1]) // р — вторая руна.

	fmt.Println(string(bytes))   // Привет — исходные байты.
	fmt.Println(string(letters)) // Привет — руны заново закодированы в UTF-8.
}

func showIteration() {
	text := "AБВ"

	// Индекс проходит по всем байтам, в том числе по обоим байтам Б и В.
	for i := 0; i < len(text); i++ {
		fmt.Println(i, text[i]) // По строкам: 0 65, 1 208, 2 145, 3 208, 4 146
	}
	fmt.Printf("%c\n", text[1]) // Ð — значение байта трактуется как кодовая точка U+00D0.

	for offset, letter := range text {
		fmt.Printf("%d: %c\n", offset, letter) // По строкам: 0: A, 1: Б, 3: В
	}

	letters := []rune(text)
	for i := 0; i < len(letters); i++ {
		fmt.Printf("%d: %c\n", i, letters[i]) // По строкам: 0: A, 1: Б, 2: В
	}

	// За некорректный байт range выдаёт U+FFFD и продвигается на один байт.
	invalid := "\xff!"
	fmt.Println(utf8.ValidString(invalid)) // false
	for offset, letter := range invalid {
		fmt.Printf("%d: %U\n", offset, letter) // По строкам: 0: U+FFFD, 1: U+0021
	}
	fmt.Printf("% X\n", string([]rune(invalid))) // EF BF BD 21 — байты руны замены и знака !.
}

func showLengths() {
	text := "Привет"
	fmt.Println(len(text))                    // 12 — байты.
	fmt.Println(len([]rune(text)))            // 6 — руны.
	fmt.Println(utf8.RuneCountInString(text)) // 6 — подсчёт без среза рун.
}

func showStringReplacement() {
	text := "Привет"
	original := text
	// text[0] = 'A' // Ошибка компиляции: байт строки не адресуем.

	letters := []rune(text)
	letters[0] = 'О'
	text = string(letters)
	fmt.Println(text, original) // Оривет Привет — original сохранила прежнюю строку.
}

func showSubstrings() {
	text := "Привет"
	fmt.Println(text[:2]) // П — граница на конце руны.
	prefix := text[:3]
	fmt.Printf("% X %t\n", prefix, utf8.ValidString(prefix)) // D0 9F D1 false — включён только первый байт р.

	letters := []rune(text)
	fmt.Println(string(letters[:3])) // При — первые три кодовые точки.
}
