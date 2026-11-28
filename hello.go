package main

import "fmt"

func main() {
	a, b := 7, 2

	fmt.Println("7 + 2 =", a+b) // 7 + 2 = 9
	fmt.Println("7 - 2 =", a-b) // 7 - 2 = 5
	fmt.Println("7 * 2 =", a*b) // 7 * 2 = 14

	// При делении целых чисел результат усекается к нулю. Оператор % даёт остаток.
	fmt.Println("7 / 2 =", a/b)                                     // 7 / 2 = 3
	fmt.Println("float64(7) / float64(2) =", float64(a)/float64(b)) // float64(7) / float64(2) = 3.5
	fmt.Println("float64(7 / 2) =", float64(a/b))                   // float64(7 / 2) = 3
	fmt.Println("7 % 2 =", a%b)                                     // 7 % 2 = 1
	fmt.Println("-7 / 2 =", -a/b)                                   // -7 / 2 = -3
	fmt.Println("-7 % 2 =", -a%b)                                   // -7 % 2 = -1

	// Умножение выполняется раньше сложения. Скобки задают другой порядок.
	fmt.Println("2 + 3 * 4 =", 2+3*4)     // 2 + 3 * 4 = 14
	fmt.Println("(2 + 3) * 4 =", (2+3)*4) // (2 + 3) * 4 = 20

	// += складывает и записывает результат в ту же переменную.
	total := 10
	total += 5
	fmt.Println("После total += 5:", total) // После total += 5: 15

	// ++ и -- изменяют значение на единицу и записываются отдельными инструкциями.
	total++
	fmt.Println("После total++:", total) // После total++: 16
	total--
	fmt.Println("После total--:", total) // После total--: 15

	// Минимальное знаковое значение при делении на -1 остаётся тем же.
	var minimum int8 = -128
	fmt.Println("Особый случай int8:", minimum/-1, minimum%-1) // Особый случай int8: -128 0

	var limit uint8 = 255
	fmt.Println("Переполнение uint8:", limit+1) // Переполнение uint8: 0
	// _ = uint8(255 + 1) // Ошибка компиляции. 256 выходит за диапазон uint8.

	// Значения в переменных float64 уже округлены до двоичной дроби.
	x, y := 0.1, 0.2
	fmt.Println("Равенство констант:", 0.1+0.2 == 0.3) // Равенство констант: true
	fmt.Println("Равенство float64:", x+y == 0.3)      // Равенство float64: false
}
