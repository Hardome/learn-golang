package main

import "fmt"

func main() {
	a, b, c := 5, 6, 5

	// Если обе проверки дали false, a меньше либо равно и b, и c.
	var relation string

	if a > b {
		relation = "a > b"
	} else if a > c {
		relation = "a > c"
	} else {
		relation = "a <= b и a <= c"
	}
	fmt.Println("Сравнение:", relation) // Сравнение: a <= b и a <= c

	if remainder := 17 % 3; remainder == 0 {
		fmt.Println("делится на 3")
	} else {
		fmt.Printf("остаток %d\n", remainder) // остаток 2
	}
	// fmt.Println(remainder) // Ошибка компиляции. Область видимости закончилась вместе с if.

	// Три значения показывают обычное совпадение, fallthrough и default.
	for _, value := range [3]int{4, 5, 7} {
		switch value {
		case 3, 4:
			fmt.Println(value, "case 3 или 4") // 4 case 3 или 4
		case 5:
			fmt.Println(value, "case 5") // 5 case 5
			fallthrough                  // Выполняет следующий блок, даже если value != 6.
		case 6:
			fmt.Println(value, "case 6") // 5 case 6
		default:
			fmt.Println(value, "default") // 7 default
		}
	}

	// Без выражения после switch каждая ветка задаёт условие.
	temperature := 18

	var weather string

	switch {
	case temperature < 0:
		weather = "морозно"
	case temperature >= 25:
		weather = "жарко"
	default:
		weather = "умеренно"
	}
	fmt.Println("На улице:", weather) // На улице: умеренно

	fmt.Println(describeType(42))   // int, следующее число 43
	fmt.Println(describeType("Go")) // строка: true
	fmt.Println(describeType(nil))  // nil
}

func describeType(value any) string {
	switch typed := value.(type) {
	case int:
		return fmt.Sprintf("int, следующее число %d", typed+1)
	case string, bool:
		_, isString := typed.(string)
		return fmt.Sprintf("строка: %t", isString)
	case nil:
		return "nil"
	default:
		return "другой тип"
	}
}
