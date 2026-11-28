package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Кавычки определяют, будет литерал строкой или руной.
	fmt.Printf("%T %T %T\n", "П", `П`, 'П') // string string int32
	// Двойные кавычки обрабатывают \n, а обратные сохраняют его как два байта.
	fmt.Printf("%q %q\n", "Go\n", `Go\n`)            // "Go\n" "Go\\n"
	fmt.Println(len("Go"), len("Привет"), len("你好")) // 2 12 6

	text := "П!"
	fmt.Println(len(text), utf8.RuneCountInString(text)) // 3 2, три байта и две руны.
	// Букву é можно записать одной руной или буквой e со знаком ударения.
	fmt.Println(utf8.RuneCountInString("é"), utf8.RuneCountInString("e\u0301")) // 1 2

	// Индекс перебирает байты, включая оба байта буквы П.
	for i := 0; i < len(text); i++ {
		fmt.Printf("%d: %02X\n", i, text[i]) // По строкам: 0: D0, 1: 9F, 2: 21
	}
	fmt.Printf("%c\n", text[0]) // Ð: %c принимает значение байта за кодовую точку U+00D0.
	// range даёт смещение первого байта руны и саму руну.
	for offset, symbol := range text {
		fmt.Printf("%d: %c\n", offset, symbol) // По строкам: 0: П, 2: !
	}
	// Строка может содержать неполный UTF-8: [:1] оставляет только байт D0.
	partial := text[:1]
	fmt.Printf("% X %t\n", partial, utf8.ValidString(partial)) // D0 false
	for offset, symbol := range partial + "!" {
		fmt.Printf("%d: %U\n", offset, symbol) // По строкам: 0: U+FFFD, 1: U+0021
	}

	// string сохраняет байты среза в том же порядке.
	received := []byte{0xD0, 0x9F, 0x21}
	fmt.Println(string(received)) // П!
	// Для замены однобайтового знака ! меняем байт и собираем новую строку.
	outgoing := "П!"
	encoded := []byte(outgoing)
	fmt.Printf("% X\n", encoded) // D0 9F 21
	encoded[2] = '?'
	outgoing = string(encoded)
	fmt.Println(outgoing) // П?

	// Числовой код руны и байты её UTF-8-представления имеют разные значения.
	var letter rune = 'П'
	var bytePairAsRune rune = 0xD09F
	fmt.Printf("%c %U %d\n", letter, letter, letter)      // П U+041F 1055
	fmt.Printf("%c %U\n", bytePairAsRune, bytePairAsRune) // 킟 U+D09F

	// В срезе рун индекс показывает порядковый номер руны от нуля.
	letters := []rune(text)
	for i := 0; i < len(letters); i++ {
		fmt.Printf("%d: %c\n", i, letters[i]) // По строкам: 0: П, 1: !
	}
	// При преобразовании среза рун в строку Go кодирует руны в UTF-8.
	fmt.Println(string([]rune{0x041F, '!'})) // П!

	// text[0] = 'A' // Ошибка компиляции: байту строки нельзя присвоить значение.
	// Срез рун позволяет заменить всю букву, даже если изменится число байтов.
	letters[0] = 'A'
	modified := string(letters)
	fmt.Println(text, modified, len(modified)) // П! A! 2
}
